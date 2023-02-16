package job

import (
	"encoding/json"
	"log"

	"github.com/dorianneto/url-shortener/src/model"
)

type CreateRedirectJob struct {
	Payload *model.Redirect
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

	// TODO: save into a database
	log.Printf(" [*] Successfully processed data from queue %s", redirect.ID)
	return nil

}
