package permissions

import (
	"example.com/PROJECT_NAME/src/core/application/response"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport/ctx"
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
