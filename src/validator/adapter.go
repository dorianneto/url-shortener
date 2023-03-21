package validator

import (
	"github.com/go-playground/validator/v10"
)

type ValidatorInterface interface {
	Struct(input interface{}) error
}

type validatorAdapter struct {
	instance *validator.Validate
}

func (v *validatorAdapter) getInstance() *validator.Validate {
	if v.instance == nil {
		v.instance = validator.New()
	}

	return v.instance
}

func (v *validatorAdapter) Struct(s interface{}) error {
	return v.getInstance().Struct(s)
}

func New() ValidatorInterface {
	return &validatorAdapter{}
}
