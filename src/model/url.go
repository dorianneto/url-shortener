package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type Redirect struct {
	ID        string    `json:"id"`
	Url       string    `json:"url" validate:"required"`
	Code      string    `json:"code" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (redirect *Redirect) isValid(validate *validator.Validate) error {
	err := validate.Struct(redirect)

	if err != nil {
		return err
	}

	return nil
}

func NewRedirect(validate *validator.Validate, url string, code string) (*Redirect, error) {
	redirect := Redirect{
		Url:  url,
		Code: code,
	}

	redirect.ID = uuid.NewV4().String()
	redirect.CreatedAt = time.Now()

	err := redirect.isValid(validate)

	if err != nil {
		return nil, err
	}

	return &redirect, nil
}
