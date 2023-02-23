package job

type JobInterface interface {
	Loader() (string, interface{})
	Handler(data interface{}) error
}
