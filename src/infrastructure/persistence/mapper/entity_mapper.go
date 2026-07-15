package mapper

import (
	"example.com/PROJECT_NAME/src/core/domain/entity"
	"example.com/PROJECT_NAME/src/infrastructure/persistence/models"
)

func EntityModelToEntity(m *models.EntityModel) *entity.EntityNameEntity {
	return &entity.EntityNameEntity{
		ID:   m.ID,
		Name: m.Name,
	}
}

func EntityEntityToModel(e *entity.EntityNameEntity) *models.EntityModel {
	return &models.EntityModel{
		Name: e.Name,
	}
}
