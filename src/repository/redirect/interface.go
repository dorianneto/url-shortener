package redirect

import (
	"github.com/dorianneto/url-shortener/src/controller/redirect/input"
	"github.com/dorianneto/url-shortener/src/model"
)

type RedirectRepositoryInterface interface {
	Find(query input.FindRedirect) (*model.Redirect, error)
	// TODO: create a output DTO with stored data
	Create(redirect *model.Redirect) (*model.Redirect, error)
}
