package repository

import (
	"github.com/dorianneto/url-shortener/src/database"
	"github.com/dorianneto/url-shortener/src/model"
)

type RedirectRepository struct {
	Database database.DatabaseInterface
}

func (rr *RedirectRepository) Find(code string) (interface{}, error) {
	result, err := rr.Database.Read(code)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rr *RedirectRepository) Create(data *model.Redirect) (interface{}, error) {
	result, err := rr.Database.Write(data.Code, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
