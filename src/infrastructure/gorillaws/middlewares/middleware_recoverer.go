package middlewares

import (
	"fmt"

	"example.com/PROJECT_NAME/src/core/domain/ports/wsport"
	"example.com/PROJECT_NAME/src/infrastructure/logger"

	"go.uber.org/zap"
)

type RecovererMiddleware struct {
	logger *logger.WsLogger
	next   wsport.MessageHandler
}

// @inject
func NewRecovererMiddleware(logger *logger.WsLogger) *RecovererMiddleware {
	return &RecovererMiddleware{logger: logger}
}

func (this *RecovererMiddleware) Wrap(next wsport.MessageHandler) wsport.MessageHandler {
	return &RecovererMiddleware{logger: this.logger, next: next}
}

func (this *RecovererMiddleware) OnOpen(ctx wsport.Context, connID wsport.ConnectionID) (err error) {
	defer func() {
		if r := recover(); r != nil {
			this.logger.Error("[recoverer] OnOpen panic", zap.String("conn_id", string(connID)), zap.Any("panic", r))
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return this.next.OnOpen(ctx, connID)
}

func (this *RecovererMiddleware) OnMessage(ctx wsport.Context, connID wsport.ConnectionID, payload []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			this.logger.Error("[recoverer] OnMessage panic", zap.String("conn_id", string(connID)), zap.Any("panic", r))
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return this.next.OnMessage(ctx, connID, payload)
}

func (this *RecovererMiddleware) OnClose(ctx wsport.Context, connID wsport.ConnectionID, closeErr error) {
	defer func() {
		if r := recover(); r != nil {
			this.logger.Error("[recoverer] OnClose panic", zap.String("conn_id", string(connID)), zap.Any("panic", r))
		}
	}()
	this.next.OnClose(ctx, connID, closeErr)
}
