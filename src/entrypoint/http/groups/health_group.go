package groups

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport"
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/http/handlers/health"
)

type HealthGroup struct {
	healthHandler *health.HealthHandler
}

// @inject
func NewHealthGroup(healthHandler *health.HealthHandler) *HealthGroup {
	return &HealthGroup{healthHandler: healthHandler}
}

func (this *HealthGroup) RegisterRoutes(g httpport.Group) {
	g.GET("", this.healthHandler.Handle)
}
