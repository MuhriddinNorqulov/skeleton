package gorillaws

import (
	"context"
	"time"

	"example.com/PROJECT_NAME/src/core/domain/ports/wsport"
)

type WsContext struct {
	context.Context
	conn *Connection
}

func newWsContext(ctx context.Context, conn *Connection) wsport.Context {
	return &WsContext{Context: ctx, conn: conn}
}

func (this *WsContext) SetReadDeadline(d time.Duration) {
	this.conn.setReadDeadline(d)
}
