package job

import (
	"encoding/json"
	"log"

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
	var redirect model.Redirect

	err := json.Unmarshal(data.([]byte), &redirect)
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := j.Repository.Create()
	if err != nil {
		log.Println(err)
	}

	log.Println(result)
	return nil

}
