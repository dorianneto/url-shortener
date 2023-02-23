package redirect

import (
	"github.com/dorianneto/url-shortener/src/controller/redirect/input"
	"github.com/dorianneto/url-shortener/src/database"
	"github.com/dorianneto/url-shortener/src/model"
)

type RedirectRepository struct {
	Database database.DocumentInterface
}

func (rr *RedirectRepository) Find(query input.FindRedirect) (*model.Redirect, error) {
	result, err := rr.Database.Read(query.Code)
	if err != nil {
		return nil, err
	}

	// TODO: narrow down access to properties externally
	redirect := &model.Redirect{Url: result.GetByKey("url")}

	return redirect, nil
}

func (rr *RedirectRepository) Create(redirect *model.Redirect) (*model.Redirect, error) {
	result, err := rr.Database.Write(redirect.Code, redirect)
	if err != nil {
		return nil, err
	}

	return result.(*model.Redirect), nil
}
