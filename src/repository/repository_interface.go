package repository

import "github.com/dorianneto/url-shortener/src/model"

type RepositoryInterface interface {
	Find(code string) (interface{}, error)
	Create(data *model.Redirect) (interface{}, error)
}
