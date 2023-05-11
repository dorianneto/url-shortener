package asynq

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dorianneto/url-shortener/src/job"
	"github.com/hibiken/asynq"
)

type asynqClientAdapterInterface interface {
	Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}

type asynqClientAdapter struct {
	client asynqClientAdapterInterface
}

func NewAsynqClientAdapter() *asynqClientAdapter {
	return &asynqClientAdapter{}
}

func (q *asynqClientAdapter) getInstance() asynqClientAdapterInterface {
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
