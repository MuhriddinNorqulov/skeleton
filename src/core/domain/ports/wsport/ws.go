package wsport

import "context"

// UpgradeFunc HTTP-dan WebSocket'ga o'tkazadi.
// Parametrlar runtime'da http.ResponseWriter va *http.Request.
// Domain layer'da net/http bo'lmasligi uchun any ishlatiladi.
type UpgradeFunc func(w any, r any)

type Ws interface {
	Init()
	Use(middlewares ...Middleware)
	Handle(channel Channel, h MessageHandler, middlewares ...Middleware)
	SendTo(ctx context.Context, connID ConnectionID, payload []byte) error
	Close(connID ConnectionID) error
	CloseAll() error
	Connect(ctx context.Context, connID ConnectionID, url string, ch Channel) error
	Upgrade(channel Channel) UpgradeFunc
}
