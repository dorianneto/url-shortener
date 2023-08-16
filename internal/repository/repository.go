package repository

import (
	"github.com/dorianneto/url-shortener/internal/controller/redirect/input"
	"github.com/dorianneto/url-shortener/internal/database"
	"github.com/dorianneto/url-shortener/internal/model"
)

type RedirectRepositoryInterface interface {
	Find(query input.FindRedirect) (*model.Redirect, error)
	// TODO: create a output DTO with stored data
	Create(redirect *model.Redirect) (*model.Redirect, error)
}

type redirectRepository struct {
	database database.DocumentInterface
}

func (rr *redirectRepository) Find(query input.FindRedirect) (*model.Redirect, error) {
	result, err := rr.database.Read(query.Code)
	if err != nil {
		return nil, err
	}

	// TODO: narrow down access to properties externally
	redirect := &model.Redirect{Url: result.GetByKey("url")}

	return redirect, nil
}

func (rr *redirectRepository) Create(redirect *model.Redirect) (*model.Redirect, error) {
	result, err := rr.database.Write(redirect.Code, redirect)
	if err != nil {
		return nil, err
	}

	return result.(*model.Redirect), nil
}

func NewRepository(database database.DocumentInterface) *redirectRepository {
	r := &redirectRepository{
		database: database,
	}

	return r
}
