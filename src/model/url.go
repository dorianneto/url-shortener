package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Redirect struct {
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

func newRedirect(url string, code string) *Redirect {
	redirect := Redirect{
		Url:  url,
		Code: code,
	}

	redirect.ID = uuid.NewV4().String()
	redirect.CreatedAt = time.Now()

	return &redirect
}
