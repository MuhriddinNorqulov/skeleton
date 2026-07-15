package httpport

import "example.com/PROJECT_NAME/src/core/domain/ports/wsport"

type Group interface {
	Group(prefix string, middlewares ...Middleware) Group
	Use(middlewares ...Middleware)
	GET(path string, handler HandlerFunc, middlewares ...Middleware)
	POST(path string, handler HandlerFunc, middlewares ...Middleware)
	PUT(path string, handler HandlerFunc, middlewares ...Middleware)
	DELETE(path string, handler HandlerFunc, middlewares ...Middleware)
	PATCH(path string, handler HandlerFunc, middlewares ...Middleware)
	Any(path string, handler HandlerFunc, middlewares ...Middleware)
	Ws(path string, handler wsport.UpgradeFunc, middlewares ...Middleware)
}

type IGroup interface {
	RegisterRoutes(g Group)
}
