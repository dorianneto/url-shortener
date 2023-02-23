package redirect

import (
	"encoding/json"
	"log"

	"github.com/dorianneto/url-shortener/src/model"
	repository "github.com/dorianneto/url-shortener/src/repository/redirect"
)

type CreateRedirectJob struct {
	Payload    *model.Redirect
	Repository repository.RedirectRepositoryInterface
}

func (j *CreateRedirectJob) queueName() string {
	return "create:redirect"
}

func (j *CreateRedirectJob) Loader() (string, interface{}) {
	return j.queueName(), j.Payload
}

func (j *CreateRedirectJob) Handler(data interface{}) error {
	var (
		redirect *model.Redirect
		err      error
	)

	err = json.Unmarshal(data.([]byte), &redirect)
	if err != nil {
		return err
	}

	redirect, err = j.Repository.Create(redirect)
	if err != nil {
		return err
	}

	log.Printf("[*] Successfully handled: %+v", redirect)

	return nil
}
