package validator

import (
	"github.com/go-playground/validator/v10"
)

type ValidatorAdapter struct {
	instance *validator.Validate
}

func (v *ValidatorAdapter) getInstance() *validator.Validate {
	if v.instance == nil {
		v.instance = validator.New()
	}

	return v.instance
}

func (v *ValidatorAdapter) Struct(s interface{}) error {
	return v.getInstance().Struct(s)
}
