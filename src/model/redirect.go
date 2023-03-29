package model

import (
	"math/rand"
	"time"

	"github.com/dorianneto/url-shortener/src/validator"
	uuid "github.com/satori/go.uuid"
)

type Redirect struct {
	ID        string    `json:"id"`
	Url       string    `json:"url" validate:"required,url"`
	Code      string    `json:"code" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (redirect *Redirect) isValid() error {
	validate := validator.New()
	err := validate.Struct(redirect)

	if err != nil {
		return err
	}

	return nil
}

func (r *Redirect) generateCode() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, 12)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func NewRedirect(url string) (*Redirect, error) {
	redirect := Redirect{
		Url: url,
	}

	redirect.ID = uuid.NewV4().String()
	redirect.Code = redirect.generateCode()
	redirect.CreatedAt = time.Now()

	err := redirect.isValid()

	if err != nil {
		return nil, err
	}

	return &redirect, nil
}
