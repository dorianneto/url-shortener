package job

import (
	"github.com/dorianneto/url-shortener/src/model"
)

type CreateRedirectInput struct {
	Data *model.Redirect
}

func (input *CreateRedirectInput) QueueName() string {
	return "create:redirect"
}
