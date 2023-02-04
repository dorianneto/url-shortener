package queue

import "github.com/dorianneto/url-shortener/src/job"

type QueueClient interface {
	Dispatch(job job.JobInterface) error
}
