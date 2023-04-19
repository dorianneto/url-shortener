package job

import (
	"encoding/json"
	"log"

	"github.com/dorianneto/url-shortener/src/model"
	"github.com/dorianneto/url-shortener/src/repository"
)

const QUEUE_NAME string = "create:redirect"

type CreateRedirectJobInterface interface {
	LoadPayload(payload interface{})
	Loader() (string, interface{})
	Handler(data []byte) error
}

type createRedirectJob struct {
	payload    *model.Redirect
	repository repository.RedirectRepositoryInterface
}

func NewCreateRedirectJob(repository repository.RedirectRepositoryInterface) *createRedirectJob {
	return &createRedirectJob{
		repository: repository,
	}
}

func (j *createRedirectJob) LoadPayload(payload interface{}) {
	j.payload = payload.(*model.Redirect)
}

func (j *createRedirectJob) Loader() (string, interface{}) {
	return QUEUE_NAME, j.payload
}

func (j *createRedirectJob) Handler(data []byte) error {
	var (
		redirect *model.Redirect
		err      error
	)

	err = json.Unmarshal(data, &redirect)
	if err != nil {
		return err
	}

	redirect, err = j.repository.Create(redirect)
	if err != nil {
		return err
	}

	log.Printf("[*] Successfully handled: %+v", redirect)

	return nil
}
