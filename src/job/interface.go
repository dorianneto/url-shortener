package job

type JobInterface interface {
	Boot() (string, interface{})
	Handler(data interface{}) error
}
