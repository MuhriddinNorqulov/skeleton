package permissions

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/application/response"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport/ctx"
)

func AuthenticatedUserPermission(next httpport.HandlerFunc) httpport.HandlerFunc {
	return func(c ctx.Context) error {
		user := c.User()
		if user == nil {
			return response.NewResponse(response.CodeUnauthorized, false, nil, "Unauthorized user")
		}
		return next(c)
	}
}
