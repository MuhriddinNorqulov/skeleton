package defaults

import (
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DeveloperBasicAuthMiddleware struct {
	env *env.Env
}

// @inject
func NewDeveloperBasicAuthMiddleware(env *env.Env) *DeveloperBasicAuthMiddleware {
	return &DeveloperBasicAuthMiddleware{env: env}
}

func (this *DeveloperBasicAuthMiddleware) Wrap(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		if username == this.env.DevUsername && password == this.env.DevPassword {
			return true, nil
		}
		return false, nil
	})(next)
}
