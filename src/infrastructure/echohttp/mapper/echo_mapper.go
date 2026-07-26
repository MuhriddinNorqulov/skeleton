package mapper

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport/ctx"

	"github.com/labstack/echo/v4"
)

func ToEchoHandler(h httpport.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(ctx.Context))
	}
}

func ToHTTPHandler(h echo.HandlerFunc) httpport.HandlerFunc {
	return func(c ctx.Context) error {
		return h(c.(echo.Context))
	}
}

func ToEchoMiddlewareMapper(middleware httpport.Middleware) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return middleware(ToHTTPHandler(next))(c.(ctx.Context))
		}
	}
}

func MiddlewareListMapper(middlewares []httpport.Middleware) []echo.MiddlewareFunc {
	mws := make([]echo.MiddlewareFunc, len(middlewares))
	for i, m := range middlewares {
		mws[i] = ToEchoMiddlewareMapper(m)
	}
	return mws
}
