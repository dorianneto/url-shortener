package queue

import "github.com/dorianneto/url-shortener/src/job"

type QueueClientInterface interface {
	Dispatch(job job.BaseJobInterface) error
}
