package mapper

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/entity"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/persistence/models"
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
