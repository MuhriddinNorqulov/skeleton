package groups

import (
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport"
	"example.com/PROJECT_NAME/src/core/domain/ports/wsport"
	wshandlers "example.com/PROJECT_NAME/src/entrypoint/http/handlers/ws"
)

type WsGroup struct {
	ws             wsport.Ws
	exampleHandler *wshandlers.ExampleWsHandler
}

// @inject
func NewWsGroup(
	ws wsport.Ws,
	exampleHandler *wshandlers.ExampleWsHandler,
) *WsGroup {
	return &WsGroup{ws: ws, exampleHandler: exampleHandler}
}

func (this *WsGroup) RegisterRoutes(g httpport.Group) {
	this.ws.Init()
	this.ws.Handle(wsport.ChannelExample, this.exampleHandler)
	g.Ws("/example", this.ws.Upgrade(wsport.ChannelExample))
}
