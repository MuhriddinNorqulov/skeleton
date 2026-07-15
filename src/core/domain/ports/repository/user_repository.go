package repository

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/entity"
	"example.com/PROJECT_NAME/src/core/domain/ports/unitofwork"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uint) (*entity.UserEntity, error)
	CreateTx(ctx context.Context, tx unitofwork.Tx, user *entity.UserEntity) (*entity.UserEntity, error)
}
