package asynq

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/dorianneto/url-shortener/internal/job"
	"github.com/hibiken/asynq"
)

type asynqServerAdapterInterface interface {
	Run(handler asynq.Handler) error
	Shutdown()
}

type asynqServerMuxAdapterInterface interface {
	HandleFunc(pattern string, handler func(context.Context, *asynq.Task) error)
	ProcessTask(context.Context, *asynq.Task) error
}

type asynqServerdapter struct {
	server  asynqServerAdapterInterface
	mux     asynqServerMuxAdapterInterface
	workers []job.BaseJobInterface
}

func NewAsynqServerdapter() *asynqServerdapter {
	return &asynqServerdapter{}
}

func (q *asynqServerdapter) getServerInstance() (asynqServerAdapterInterface, error) {
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

func (q *asynqServerdapter) getMuxInstance() asynqServerMuxAdapterInterface {
	if q.mux == nil {
		q.mux = asynq.NewServeMux()
	}

	return q.mux
}

func (q *asynqServerdapter) RegisterWorker(handler job.BaseJobInterface) {
	q.workers = append(q.workers, handler)
}

func (q *asynqServerdapter) RunWorkers() {
	server, err := q.getServerInstance()
	if err != nil {
		log.Println(err)
		return
	}

	mux := q.getMuxInstance()

	for _, w := range q.workers {
		queue, _ := w.Loader()

		mux.HandleFunc(queue, func(c context.Context, t *asynq.Task) error {
			err := w.Handler(t.Payload())

			if err != nil {
				log.Println(err)
			}

			return err
		})
	}

	log.Println("Redis server running...")

	if err := server.Run(mux.(asynq.Handler)); err != nil {
		server.Shutdown()
		log.Println(err)
		return
	}
}
