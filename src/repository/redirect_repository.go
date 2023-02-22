package repository

import (
	"github.com/dorianneto/url-shortener/src/database"
	"github.com/dorianneto/url-shortener/src/model"
)

type RedirectRepository struct {
	Database database.DocumentInterface
}

func (rr *RedirectRepository) Find(code string) (interface{}, error) {
	result, err := rr.Database.Read(code)
	if err != nil {
		return nil, err
	}

	redirect := &model.Redirect{Url: result.Data["Url"].(string)}

	return redirect, nil
}

func (rr *RedirectRepository) Create(data interface{}) (interface{}, error) {
	redirect := data.(*model.Redirect)

	result, err := rr.Database.Write(redirect.Code, redirect)
	if err != nil {
		return nil, err
	}

	return result, nil
}
