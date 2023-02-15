package job

type JobInterface interface {
	PreEnqueue() InputInterface
	// Handle() error
}
