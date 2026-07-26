package gorillaws

import (
	"context"
	"sync"
	"time"

	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"

	"github.com/gorilla/websocket"
)

type Connection struct {
	id      wsport.ConnectionID
	channel wsport.Channel
	ws      *websocket.Conn
	writeMu sync.Mutex
	cancel  context.CancelFunc
}

func newConnection(url string, topic wsport.Channel, id wsport.ConnectionID, ws *websocket.Conn, cancel context.CancelFunc) *Connection {
	return &Connection{id: id, channel: topic, ws: ws, cancel: cancel}
}

func (this *Connection) reset(ws *websocket.Conn) {
	this.ws = ws
}

func (this *Connection) setReadDeadline(d time.Duration) {
	_ = this.ws.SetReadDeadline(time.Now().Add(d))
}

func (this *Connection) read() ([]byte, error) {
	_, payload, err := this.ws.ReadMessage()
	return payload, err
}

func (this *Connection) send(payload []byte) error {
	this.writeMu.Lock()
	defer this.writeMu.Unlock()
	return this.ws.WriteMessage(websocket.TextMessage, payload)
}

func (this *Connection) close() error {
	if this.cancel != nil {
		this.cancel()
	}
	return this.ws.Close()
}
