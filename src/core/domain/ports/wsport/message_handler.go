package wsport

type MessageHandler interface {
	OnOpen(ctx Context, connID ConnectionID) error
	OnMessage(ctx Context, connID ConnectionID, payload []byte) error
	OnClose(ctx Context, connID ConnectionID, err error)
}
