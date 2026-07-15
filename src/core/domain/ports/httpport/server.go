package httpport

type HTTPServer interface {
	Use(middlewares ...Middleware)
	Run()
	Group(prefix string, middlewares ...Middleware) Group
	Init()
}
