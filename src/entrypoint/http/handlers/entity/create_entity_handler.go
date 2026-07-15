package entity

import (
	"example.com/PROJECT_NAME/src/core/application/dto"
	"example.com/PROJECT_NAME/src/core/application/usecases/entityusecases"
	entitypkg "example.com/PROJECT_NAME/src/core/domain/entity"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport/ctx"
	"net/http"
)

type CreateEntityHandler struct {
	uc *entityusecases.CreateEntityUseCase
}

// @inject
func NewCreateEntityHandler(uc *entityusecases.CreateEntityUseCase) *CreateEntityHandler {
	return &CreateEntityHandler{uc: uc}
}

// Handle godoc
// @Tags         Entity
// @Summary      Create entity
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateEntityRequest  true  "Entity data"
// @Success      200   {object}  response.Response{payload=dto.CreateEntityResult}
// @Security     BearerAuth
// @Router       /v1/entity/create [post]
func (this *CreateEntityHandler) Handle(c ctx.Context) error {
	req, err := ctx.GetBody[dto.CreateEntityRequest](c)
	if err != nil {
		return err
	}

	id, err := this.uc.Invoke(c.GetContext(), &entitypkg.EntityNameEntity{Name: req.Name})
	if err != nil {
		return err
	}

	return c.Success(http.StatusOK, &dto.CreateEntityResult{ID: id})
}
