package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	FirstName string  `gorm:"size:64;"`
	LastName  *string `gorm:"size:64;"`
}

func (this *UserModel) TableName() string {
	return "users"
}
