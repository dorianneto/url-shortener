package job

type BaseJobInterface interface {
	LoadPayload(payload interface{})
	Loader() (string, interface{})
	Handler(data []byte) error
}
