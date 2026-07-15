package defaults

import (
	"example.com/PROJECT_NAME/src/infrastructure/logger"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RecoveryMiddleware struct {
	logger *logger.HttpLogger
}

// @inject
func NewRecoveryMiddleware(logger *logger.HttpLogger) *RecoveryMiddleware {
	return &RecoveryMiddleware{logger: logger}
}

func (this *RecoveryMiddleware) Wrap(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				stack := make([]byte, 4<<10)
				length := runtime.Stack(stack, false)

				req := c.Request()

				this.logger.Error("panic recovered",
					zap.Any("panic", r),
					zap.ByteString("stack", stack[:length]),
					zap.String("method", req.Method),
					zap.String("path", req.URL.Path),
					zap.String("ip", c.RealIP()),
				)

				err = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
			}
		}()

		return next(c)
	}
}
