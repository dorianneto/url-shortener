package createredirectjob

import (
	"github.com/dorianneto/url-shortener/src/job"
	"github.com/dorianneto/url-shortener/src/model"
)

type CreateRedirectJob struct {
	Data *model.Redirect
}

func (job *CreateRedirectJob) PreEnqueue() job.BaseInputInterface {
	return &CreateRedirectInput{Data: job.Data}
}
