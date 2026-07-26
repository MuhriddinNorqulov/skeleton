package repository

import (
	"context"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/entity"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/unitofwork"
)

type EntityRepository interface {
	GetByID(ctx context.Context, id uint) (*entity.EntityNameEntity, error)
	Create(ctx context.Context, e *entity.EntityNameEntity) (uint, error)
	CreateTx(ctx context.Context, tx unitofwork.Tx, e *entity.EntityNameEntity) (uint, error)
}
