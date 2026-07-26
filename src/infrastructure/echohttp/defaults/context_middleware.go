package defaults

import (
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/echohttp/context"

	"github.com/labstack/echo/v4"
)

func ContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.NewEchoContext(c)
		return next(ctx.(echo.Context))
	}
}
