package jobtest

var LoaderFnDefaultValue = func() (string, interface{}) {
	return "foo", map[string]int{"foo": 1, "bar": 2, "baz": 3}
}

var LoaderFn = LoaderFnDefaultValue

type JobMock struct{}

func (j *JobMock) LoadPayload(payload interface{}) {}

func (j *JobMock) Loader() (string, interface{}) {
	return LoaderFn()
}

func (j *JobMock) Handler(data []byte) error {
	return nil
}
