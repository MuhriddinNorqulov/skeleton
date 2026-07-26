package groups

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"
	wshandlers "github.com/muhriddinnorqulov/skeleton/src/entrypoint/http/handlers/ws"
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
