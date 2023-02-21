package job

import (
	"encoding/json"

	"github.com/dorianneto/url-shortener/src/model"
	"github.com/dorianneto/url-shortener/src/repository"
)

type CreateRedirectJob struct {
	Payload    *model.Redirect
	Repository repository.RepositoryInterface
}

func (j *CreateRedirectJob) queueName() string {
	return "create:redirect"
}

func (j *CreateRedirectJob) Boot() (string, interface{}) {
	return j.queueName(), j.Payload
}

func (j *CreateRedirectJob) Handler(data interface{}) error {
	var (
		err      error
		redirect *model.Redirect
	)

	err = json.Unmarshal(data.([]byte), &redirect)
	if err != nil {
		return err
	}

	_, err = j.Repository.Create(redirect)
	if err != nil {
		return err
	}

	return nil
}
