package groups

import (
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport"
	"example.com/PROJECT_NAME/src/entrypoint/http/handlers/entity"
	"example.com/PROJECT_NAME/src/entrypoint/http/interceptor/permissions"
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
