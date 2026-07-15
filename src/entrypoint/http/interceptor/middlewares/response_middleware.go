package middlewares

import (
	"errors"
	"example.com/PROJECT_NAME/src/core/application/response"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport/ctx"
)

type ResponseMiddleware struct{}

// @inject
func NewResponseMiddleware() *ResponseMiddleware {
	return &ResponseMiddleware{}
}

func (this *ResponseMiddleware) Wrap(next httpport.HandlerFunc) httpport.HandlerFunc {
	return func(c ctx.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		if resp, ok := errors.AsType[*response.Response](err); ok {
			return c.JSON(this.codeToStatus(resp.Code), resp)
		}

		return err
	}
}

func (this *ResponseMiddleware) codeToStatus(c response.Code) int {
	switch c {
	case response.CodeBadRequest:
		return 400
	case response.CodeUnauthorized, response.CodeExpiredToken, response.CodeInvalidToken:
		return 401
	case response.CodeForbidden:
		return 403
	case response.CodeDatabaseError, response.CodeGatewayError:
		return 500
	}
	return 200
}
