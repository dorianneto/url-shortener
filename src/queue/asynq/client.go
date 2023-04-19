package queue

import (
	"log"
	"os"

	"github.com/dorianneto/url-shortener/src/job"
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
)

type AsynqClientAdapter struct {
	client *asynq.Client
}

func (q *AsynqClientAdapter) getInstance() *asynq.Client {
	if q.client == nil {
		q.client = asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")})
	}

	return q.client
}

func (q *AsynqClientAdapter) Dispatch(job job.BaseJobInterface) error {
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
