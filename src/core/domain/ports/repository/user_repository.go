package repository

import (
	"context"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/entity"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/unitofwork"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uint) (*entity.UserEntity, error)
	CreateTx(ctx context.Context, tx unitofwork.Tx, user *entity.UserEntity) (*entity.UserEntity, error)
}
