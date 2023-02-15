package model

import (
	"time"

	"github.com/dorianneto/url-shortener/src/validator"
	uuid "github.com/satori/go.uuid"
)

type Redirect struct {
	ID        string    `json:"id"`
	Url       string    `json:"url" validate:"required"`
	Code      string    `json:"code" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (redirect *Redirect) isValid() error {
	validate := validator.ValidatorAdapter{}
	err := validate.Struct(redirect)

	if err != nil {
		return err
	}

	return nil
}

func NewRedirect(url string, code string) (*Redirect, error) {
	redirect := Redirect{
		Url:  url,
		Code: code,
	}

	redirect.ID = uuid.NewV4().String()
	redirect.CreatedAt = time.Now()

	err := redirect.isValid()

	if err != nil {
		return nil, err
	}

	return &redirect, nil
}
