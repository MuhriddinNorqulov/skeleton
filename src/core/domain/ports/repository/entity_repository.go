package repository

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/entity"
	"example.com/PROJECT_NAME/src/core/domain/ports/unitofwork"
)

type EntityRepository interface {
	GetByID(ctx context.Context, id uint) (*entity.EntityNameEntity, error)
	Create(ctx context.Context, e *entity.EntityNameEntity) (uint, error)
	CreateTx(ctx context.Context, tx unitofwork.Tx, e *entity.EntityNameEntity) (uint, error)
}
