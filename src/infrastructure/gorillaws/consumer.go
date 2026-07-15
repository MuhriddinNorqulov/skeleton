package gorillaws

import "example.com/PROJECT_NAME/src/core/domain/ports/wsport"

type Consumer struct {
	mux *Mux
}

// @inject
func NewConsumer(mux *Mux) *Consumer {
	return &Consumer{mux: mux}
}

// Run — connection lifecycle: OnOpen -> message loop -> OnClose
func (this *Consumer) Run(ctx wsport.Context, c *Connection) {
	h, ok := this.mux.Get(c.channel)
	if !ok {
		return
	}

	if err := h.OnOpen(ctx, c.id); err != nil {
		return
	}

	for {
		payload, err := c.read()
		if err != nil {
			h.OnClose(ctx, c.id, err)
			return
		}
		_ = h.OnMessage(ctx, c.id, payload)
	}
}
