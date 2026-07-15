package defaults

import (
	"example.com/PROJECT_NAME/src/core/application/response"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport"

	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	validator *validator.Validate
}

// @inject
func NewRequestValidator() *RequestValidator {
	v := validator.New()
	return &RequestValidator{validator: v}
}

func (this *RequestValidator) Validate(i interface{}) error {
	if err := this.validator.Struct(i); err != nil {
		return response.NewResponse(response.CodeBadRequest, false, nil, err.Error())
	}
	if v, ok := i.(httpport.Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
