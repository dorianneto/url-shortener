package job

type JobInterface interface {
	Loader() (string, interface{})
	Handler(data []byte) error
}
