package queue

import "github.com/dorianneto/url-shortener/src/job"

type QueueClientInterface interface {
	Dispatch(job job.BaseJobInterface) error
}

type QueueServerInterface interface {
	RunWorkers() error
	RegisterWorker(handler job.BaseJobInterface) error
}
