package ws

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"
)

type ExampleWsHandler struct {
	// TODO: use case dependency
}

// @inject
func NewExampleWsHandler() *ExampleWsHandler {
	return &ExampleWsHandler{}
}

func (this *ExampleWsHandler) OnOpen(_ wsport.Context, _ wsport.ConnectionID) error {
	// TODO: ulanish ochilgandagi logika
	return nil
}

func (this *ExampleWsHandler) OnMessage(_ wsport.Context, _ wsport.ConnectionID, _ []byte) error {
	// TODO: xabar kelgandagi logika
	return nil
}

func (this *ExampleWsHandler) OnClose(_ wsport.Context, _ wsport.ConnectionID, _ error) {
	// TODO: ulanish yopilgandagi tozalash
}
