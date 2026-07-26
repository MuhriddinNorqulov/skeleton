package echohttp

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/echohttp/mapper"

	"github.com/labstack/echo/v4"
)

type EchoGroupImpl struct {
	g *echo.Group
}

func NewEchoGroupImpl(g *echo.Group) httpport.Group {
	return &EchoGroupImpl{g: g}
}

func (this *EchoGroupImpl) Group(prefix string, middlewares ...httpport.Middleware) httpport.Group {
	return NewEchoGroupImpl(this.g.Group(prefix, mapper.MiddlewareListMapper(middlewares)...))
}

func (this *EchoGroupImpl) Use(middlewares ...httpport.Middleware) {
	this.g.Use(mapper.MiddlewareListMapper(middlewares)...)
}

func (this *EchoGroupImpl) GET(path string, handler httpport.HandlerFunc, middlewares ...httpport.Middleware) {
	this.g.GET(path, mapper.ToEchoHandler(handler), mapper.MiddlewareListMapper(middlewares)...)
}

func (this *EchoGroupImpl) POST(path string, handler httpport.HandlerFunc, middlewares ...httpport.Middleware) {
	this.g.POST(path, mapper.ToEchoHandler(handler), mapper.MiddlewareListMapper(middlewares)...)
}

func (this *EchoGroupImpl) PUT(path string, handler httpport.HandlerFunc, middlewares ...httpport.Middleware) {
	this.g.PUT(path, mapper.ToEchoHandler(handler), mapper.MiddlewareListMapper(middlewares)...)
}

func (this *EchoGroupImpl) DELETE(path string, handler httpport.HandlerFunc, middlewares ...httpport.Middleware) {
	this.g.DELETE(path, mapper.ToEchoHandler(handler), mapper.MiddlewareListMapper(middlewares)...)
}

func (this *EchoGroupImpl) PATCH(path string, handler httpport.HandlerFunc, middlewares ...httpport.Middleware) {
	this.g.PATCH(path, mapper.ToEchoHandler(handler), mapper.MiddlewareListMapper(middlewares)...)
}

func (this *EchoGroupImpl) Any(path string, handler httpport.HandlerFunc, middlewares ...httpport.Middleware) {
	this.g.Any(path, mapper.ToEchoHandler(handler), mapper.MiddlewareListMapper(middlewares)...)
}

func (this *EchoGroupImpl) Ws(path string, handler wsport.UpgradeFunc, middlewares ...httpport.Middleware) {
	this.g.GET(path, func(c echo.Context) error {
		handler(c.Response().Writer, c.Request())
		return nil
	}, mapper.MiddlewareListMapper(middlewares)...)
}
