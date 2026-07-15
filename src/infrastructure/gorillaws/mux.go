package gorillaws

import (
	"sync"

	"example.com/PROJECT_NAME/src/core/domain/ports/wsport"
)

type Mux struct {
	mu                sync.RWMutex
	handlers          map[wsport.Channel]wsport.MessageHandler
	globalMiddlewares []wsport.Middleware
}

// @inject
func NewMux() *Mux {
	return &Mux{handlers: make(map[wsport.Channel]wsport.MessageHandler)}
}

func (this *Mux) Use(middlewares ...wsport.Middleware) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.globalMiddlewares = append(this.globalMiddlewares, middlewares...)
}

func (this *Mux) Register(topic wsport.Channel, h wsport.MessageHandler, middlewares ...wsport.Middleware) {
	all := make([]wsport.Middleware, 0, len(this.globalMiddlewares)+len(middlewares))
	all = append(all, this.globalMiddlewares...)
	all = append(all, middlewares...)

	this.mu.Lock()
	defer this.mu.Unlock()
	this.handlers[topic] = wsport.Chain(h, all...)
}

func (this *Mux) Get(topic wsport.Channel) (wsport.MessageHandler, bool) {
	this.mu.RLock()
	defer this.mu.RUnlock()
	h, ok := this.handlers[topic]
	return h, ok
}
