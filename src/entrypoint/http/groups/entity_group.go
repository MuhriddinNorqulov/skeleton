package groups

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport"
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/http/handlers/entity"
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/http/interceptor/permissions"
)

type EntityGroup struct {
	createHandler *entity.CreateEntityHandler
}

// @inject
func NewEntityGroup(createHandler *entity.CreateEntityHandler) *EntityGroup {
	return &EntityGroup{createHandler: createHandler}
}

func (this *EntityGroup) RegisterRoutes(g httpport.Group) {
	g.POST("/create", this.createHandler.Handle, permissions.AuthenticatedUserPermission)
}
