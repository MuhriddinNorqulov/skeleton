package wsport

type Middleware interface {
	Wrap(next MessageHandler) MessageHandler
}

func Chain(h MessageHandler, middlewares ...Middleware) MessageHandler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i].Wrap(h)
	}
	return h
}
