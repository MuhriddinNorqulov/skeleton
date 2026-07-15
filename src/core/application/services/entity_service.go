package services

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/ports/repository"
	"example.com/PROJECT_NAME/src/core/domain/ports/unitofwork"
)

type EntityService struct {
	repo   repository.EntityRepository
	atomic unitofwork.Atomic
}

// @inject
func NewEntityService(
	repo repository.EntityRepository,
	atomic unitofwork.Atomic,
) *EntityService {
	return &EntityService{repo: repo, atomic: atomic}
}

func (this *EntityService) DoSomething(ctx context.Context, id uint) error {
	// TODO: qayta ishlatiladigan biznes logika
	_, err := this.repo.GetByID(ctx, id)
	return err
}
