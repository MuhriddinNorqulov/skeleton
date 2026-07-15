package repository

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/entity"
	"example.com/PROJECT_NAME/src/core/domain/ports/repository"
	"example.com/PROJECT_NAME/src/core/domain/ports/unitofwork"
	"example.com/PROJECT_NAME/src/infrastructure/_errors"
	"example.com/PROJECT_NAME/src/infrastructure/persistence/mapper"
	"example.com/PROJECT_NAME/src/infrastructure/persistence/models"
)

type EntityRepositoryImpl struct {
	*BaseRepository
}

// @inject
func NewEntityRepositoryImpl(base *BaseRepository) repository.EntityRepository {
	return &EntityRepositoryImpl{BaseRepository: base}
}

func (this *EntityRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.EntityNameEntity, error) {
	var model models.EntityModel
	if err := this.db(ctx).First(&model, id).Error; err != nil {
		return nil, _errors.GormErrorWrap(err)
	}
	return mapper.EntityModelToEntity(&model), nil
}

func (this *EntityRepositoryImpl) Create(ctx context.Context, e *entity.EntityNameEntity) (uint, error) {
	model := mapper.EntityEntityToModel(e)
	if err := this.db(ctx).Create(model).Error; err != nil {
		return 0, _errors.GormErrorWrap(err)
	}
	return model.ID, nil
}

func (this *EntityRepositoryImpl) CreateTx(ctx context.Context, tx unitofwork.Tx, e *entity.EntityNameEntity) (uint, error) {
	model := mapper.EntityEntityToModel(e)
	if err := this.tx(ctx, tx).Create(model).Error; err != nil {
		return 0, _errors.GormErrorWrap(err)
	}
	return model.ID, nil
}
