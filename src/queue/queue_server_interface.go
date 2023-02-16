package queue

import "github.com/dorianneto/url-shortener/src/job"

type QueueServerInterface interface {
	RunWorkers() error
	RegisterWorker(handler job.JobInterface) error
}
