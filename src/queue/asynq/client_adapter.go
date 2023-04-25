package asynq

import (
	"log"
	"os"

	"github.com/dorianneto/url-shortener/src/job"
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
)

type asynqClientAdapter struct {
	client *asynq.Client
}

func NewAsynqClientAdapter() *asynqClientAdapter {
	return &asynqClientAdapter{}
}

func (q *asynqClientAdapter) getInstance() *asynq.Client {
	if q.client == nil {
		q.client = asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")})
	}

	return q.client
}

func (q *asynqClientAdapter) Dispatch(job job.BaseJobInterface) error {
	queue, data := job.Loader()

	dataEncoded, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := q.getInstance()

	result, err := client.Enqueue(asynq.NewTask(queue, dataEncoded))
	if err != nil {
		return err
	}

	log.Printf("[*] Successfully enqueued task: %+v", result)

	return nil
}
