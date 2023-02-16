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

func (queue *AsynqClientAdapter) getInstance() *asynq.Client {
	if queue.client == nil {
		queue.client = asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")})
	}

	return queue.client
}

func (queue *AsynqClientAdapter) Dispatch(job job.JobInterface) error {
	typeName, data := job.Boot()

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := queue.getInstance()

	result, err := client.Enqueue(asynq.NewTask(typeName, payload))
	if err != nil {
		return err
	}

	log.Printf("[*] Successfully enqueued task: %+v", result)

	return nil
}
