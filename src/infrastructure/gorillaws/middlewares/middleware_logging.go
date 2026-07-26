package middlewares

import (
	"time"

	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/logger"

	"go.uber.org/zap"
)

type LoggingMiddleware struct {
	logger *logger.WsLogger
	next   wsport.MessageHandler
}

// @inject
func NewLoggingMiddleware(logger *logger.WsLogger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

func (this *LoggingMiddleware) Wrap(next wsport.MessageHandler) wsport.MessageHandler {
	return &LoggingMiddleware{logger: this.logger, next: next}
}

func (this *LoggingMiddleware) OnOpen(ctx wsport.Context, connID wsport.ConnectionID) error {
	start := time.Now()
	err := this.next.OnOpen(ctx, connID)
	if err != nil {
		this.logger.Error("OnOpen error", zap.String("conn_id", string(connID)), zap.Error(err))
	} else {
		this.logger.Info("OnOpen", zap.String("conn_id", string(connID)), zap.Duration("duration", time.Since(start)))
	}
	return err
}

func (this *LoggingMiddleware) OnMessage(ctx wsport.Context, connID wsport.ConnectionID, payload []byte) error {
	err := this.next.OnMessage(ctx, connID, payload)
	if err != nil {
		this.logger.Error("OnMessage error", zap.String("conn_id", string(connID)), zap.Error(err))
	}
	return err
}

func (this *LoggingMiddleware) OnClose(ctx wsport.Context, connID wsport.ConnectionID, closeErr error) {
	this.next.OnClose(ctx, connID, closeErr)
}
