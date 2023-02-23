package queue

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/dorianneto/url-shortener/src/job"
	"github.com/hibiken/asynq"
)

type AsynqServerdapter struct {
	server  *asynq.Server
	mux     *asynq.ServeMux
	workers []job.JobInterface
}

func (q *AsynqServerdapter) getServerInstance() (*asynq.Server, error) {
	if q.server == nil {
		concurrency, err := strconv.Atoi(os.Getenv("REDIS_CONCURRENCY"))
		if err != nil {
			return nil, errors.New("concurrency is mandatory to start Redis workers")
		}

		q.server = asynq.NewServer(
			asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")},
			asynq.Config{Concurrency: concurrency},
		)
	}

	return q.server, nil
}

func (q *AsynqServerdapter) getMuxInstance() *asynq.ServeMux {
	if q.mux == nil {
		q.mux = asynq.NewServeMux()
	}

	return q.mux
}

func (q *AsynqServerdapter) RegisterWorker(handler job.JobInterface) {
	q.workers = append(q.workers, handler)
}

func (q *AsynqServerdapter) RunWorkers() {
	server, err := q.getServerInstance()
	if err != nil {
		log.Fatal(err)
	}

	mux := q.getMuxInstance()

	for _, w := range q.workers {
		queue, _ := w.Loader()

		mux.HandleFunc(queue, func(c context.Context, t *asynq.Task) error {
			return w.Handler(t.Payload())
		})
	}

	log.Println("Redis server running...")

	if err := server.Run(mux); err != nil {
		server.Shutdown()
		log.Fatal(err)
	}
}
