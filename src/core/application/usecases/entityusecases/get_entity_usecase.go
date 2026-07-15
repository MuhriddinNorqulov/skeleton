package entityusecases

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/entity"
	"example.com/PROJECT_NAME/src/core/domain/ports/repository"
)

type GetEntityUseCase struct {
	repo repository.EntityRepository
}

// @inject
func NewGetEntityUseCase(repo repository.EntityRepository) *GetEntityUseCase {
	return &GetEntityUseCase{repo: repo}
}

func (this *GetEntityUseCase) Invoke(ctx context.Context, id uint) (*entity.EntityNameEntity, error) {
	return this.repo.GetByID(ctx, id)
}
