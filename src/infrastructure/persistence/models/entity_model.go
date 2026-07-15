package models

import "gorm.io/gorm"

type EntityModel struct {
	gorm.Model
	Name   string `gorm:"size:256;not null;"`
	Status string `gorm:"size:32;not null;default:'PENDING';"`
}

func (this *EntityModel) TableName() string {
	return "entities"
}
