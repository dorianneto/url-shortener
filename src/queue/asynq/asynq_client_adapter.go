package queue

import (
	"log"

	"github.com/dorianneto/url-shortener/src/job"
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
)

type AsynqClientAdapter struct {
	client *asynq.Client
}

func (queue *AsynqClientAdapter) getInstance() *asynq.Client {
	if queue.client == nil {
		queue.client = asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	}

	return queue.client
}

func (queue *AsynqClientAdapter) Dispatch(job job.JobInterface) error {
	input := job.PreEnqueue()

	typeName := input.QueueName()
	payload, err := json.Marshal(input)

	if err != nil {
		return err
	}

	result, err := queue.getInstance().Enqueue(asynq.NewTask(typeName, payload))

	if err != nil {
		return err
	}

	log.Printf(" [*] Successfully enqueued task: %+v", result)

	return nil
}
