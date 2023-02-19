package repository

import "github.com/dorianneto/url-shortener/src/database"

type RedirectRepository struct {
	Database database.DatabaseInterface
}

func (rr *RedirectRepository) Find() (interface{}, error) {
	result, err := rr.Database.Read()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rr *RedirectRepository) Create() (interface{}, error) {
	result, err := rr.Database.Write()
	if err != nil {
		return nil, err
	}

	return result, nil
}
