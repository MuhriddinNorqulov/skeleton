package entityusecases

import (
	"context"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/entity"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/repository"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/unitofwork"
)

type CreateEntityUseCase struct {
	atomic unitofwork.Atomic
	repo   repository.EntityRepository
}

// @inject
func NewCreateEntityUseCase(
	atomic unitofwork.Atomic,
	repo repository.EntityRepository,
) *CreateEntityUseCase {
	return &CreateEntityUseCase{atomic: atomic, repo: repo}
}

func (this *CreateEntityUseCase) Invoke(ctx context.Context, e *entity.EntityNameEntity) (uint, error) {
	var id uint
	if err := this.atomic.Transaction(func(tx unitofwork.Tx) error {
		var err error
		id, err = this.repo.CreateTx(ctx, tx, e)
		return err
	}); err != nil {
		return 0, err
	}
	return id, nil
}
