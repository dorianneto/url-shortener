package job

type JobInterface interface {
	PreEnqueue() BaseInputInterface
	// Handle() error
}
