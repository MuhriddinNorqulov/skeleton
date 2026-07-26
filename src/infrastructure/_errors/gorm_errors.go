package _errors

import (
	"errors"
	"github.com/muhriddinnorqulov/skeleton/src/core/application/response"
	"github.com/muhriddinnorqulov/skeleton/src/core/utils"

	"gorm.io/gorm"
)

func GormErrorWrap(err error) error {
	if err == nil {
		return nil
	}

	caller := utils.CallerPath(2)

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return response.NewSafeError(response.CodeNotFound, err, caller)
	}

	return response.NewSafeError(response.CodeDatabaseError, err, caller)
}
