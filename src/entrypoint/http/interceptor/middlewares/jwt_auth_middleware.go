package middlewares

import (
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport/ctx"
	"strings"
)

type JwtAuthMiddleware struct {
	// TODO: auth service dependency
}

// @inject
func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func (this *JwtAuthMiddleware) Wrap(next httpport.HandlerFunc) httpport.HandlerFunc {
	return func(c ctx.Context) error {
		authHeader := c.GetRequest().Header.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			// TODO: token ni tekshirib, user ni set qilish
			// token := strings.TrimPrefix(authHeader, "Bearer ")
			// user, err := this.authService.VerifyToken(token)
			// if err == nil { c.SetUser(user) }
		}
		return next(c)
	}
}
