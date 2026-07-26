package health

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport/ctx"
	"net/http"
)

type HealthHandler struct{}

// @inject
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Handle godoc
// @Tags         Health
// @Summary      Health check
// @Produce      json
// @Success      200  {object}  response.Response
// @Router       /v1/health [get]
func (this *HealthHandler) Handle(c ctx.Context) error {
	return c.Success(http.StatusOK, map[string]string{"status": "ok"})
}
