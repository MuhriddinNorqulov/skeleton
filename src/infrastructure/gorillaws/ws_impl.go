package gorillaws

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/gorillaws/middlewares"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/logger"

	"github.com/google/uuid"
)

type WsImpl struct {
	upgrader            *GorillaUpgrader
	dialer              *Dialer
	registry            *Mux
	store               *ConnectionStore
	consumer            *Consumer
	logger              *logger.WsLogger
	loggingMiddleware   *middlewares.LoggingMiddleware
	recovererMiddleware *middlewares.RecovererMiddleware
}

// @inject
func NewWsImpl(
	upgrader *GorillaUpgrader,
	dialer *Dialer,
	store *ConnectionStore,
	mux *Mux,
	consumer *Consumer,
	log *logger.WsLogger,
	loggingMiddleware *middlewares.LoggingMiddleware,
	recovererMiddleware *middlewares.RecovererMiddleware,
) wsport.Ws {
	return &WsImpl{
		upgrader:            upgrader,
		dialer:              dialer,
		registry:            mux,
		store:               store,
		consumer:            consumer,
		logger:              log,
		loggingMiddleware:   loggingMiddleware,
		recovererMiddleware: recovererMiddleware,
	}
}

func (this *WsImpl) Init() {
	this.registry.Use(
		this.recovererMiddleware,
		this.loggingMiddleware,
	)
}

func (this *WsImpl) Use(mws ...wsport.Middleware) {
	this.registry.Use(mws...)
}

func (this *WsImpl) Handle(channel wsport.Channel, h wsport.MessageHandler, mws ...wsport.Middleware) {
	this.registry.Register(channel, h, mws...)
}

func (this *WsImpl) SendTo(_ context.Context, connID wsport.ConnectionID, payload []byte) error {
	conn, ok := this.store.Get(connID)
	if !ok {
		return errors.New("connection topilmadi: " + string(connID))
	}
	this.logger.MessageSent(string(connID), payload)
	return conn.send(payload)
}

func (this *WsImpl) Close(connID wsport.ConnectionID) error {
	conn, ok := this.store.Get(connID)
	if !ok {
		return nil
	}
	this.store.Remove(connID)
	return conn.close()
}

func (this *WsImpl) CloseAll() error {
	for _, conn := range this.store.All() {
		_ = conn.close()
	}
	return nil
}

func (this *WsImpl) Connect(ctx context.Context, connID wsport.ConnectionID, url string, channel wsport.Channel) error {
	ws, err := this.dialer.Dial(url)
	if err != nil {
		return err
	}

	internalCtx, cancel := context.WithCancel(ctx) //nolint:contextcheck
	conn := newConnection(url, channel, connID, ws, cancel)
	this.store.Add(conn)

	this.logger.Connected(string(connID), string(channel), url)

	go this.runWithReconnect(newWsContext(internalCtx, conn), conn, url)

	return nil
}

func (this *WsImpl) Upgrade(channel wsport.Channel) wsport.UpgradeFunc {
	return func(w any, r any) {
		httpR := r.(*http.Request)
		ws, err := this.upgrader.Upgrade(w.(http.ResponseWriter), httpR)
		if err != nil {
			this.logger.UpgradeError(string(channel), httpR.RemoteAddr, err)
			return
		}

		connID := this.generateID()
		conn := newConnection(httpR.RemoteAddr, channel, connID, ws, nil)
		this.store.Add(conn)
		defer this.store.Remove(connID)

		this.consumer.Run(newWsContext(context.WithoutCancel(httpR.Context()), conn), conn)
	}
}

func (this *WsImpl) runWithReconnect(ctx wsport.Context, conn *Connection, url string) {
	attempt := 0

	for {
		this.consumer.Run(ctx, conn)

		this.store.Remove(conn.id)

		if ctx.Err() != nil {
			return
		}

		delay := this.dialer.BackoffDelay(attempt)
		attempt++

		select {
		case <-ctx.Done():
			return
		case <-time.After(delay):
		}

		ws, err := this.dialer.Dial(url)
		if err != nil {
			this.logger.DialError(url, string(conn.channel), err)
			continue
		}

		conn.reset(ws)
		this.store.Add(conn)
		attempt = 0

		this.logger.Connected(string(conn.id), string(conn.channel), url)
	}
}

func (this *WsImpl) generateID() wsport.ConnectionID {
	return wsport.ConnectionID(fmt.Sprintf("conn-%s", uuid.NewString()))
}
