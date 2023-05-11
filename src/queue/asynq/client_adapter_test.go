package asynq

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hibiken/asynq"
)

type clientMock struct{}

var enqueueFnDefaultValue = func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return nil, nil
}
var enqueueFn = enqueueFnDefaultValue

func (cm *clientMock) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return enqueueFn(task, opts...)
}

type jobMock struct{}

var loaderFnDefaultValue = func() (string, interface{}) {
	return "foo", map[string]int{"foo": 1, "bar": 2, "baz": 3}
}
var loaderFn = loaderFnDefaultValue

func (j *jobMock) LoadPayload(payload interface{}) {}

func (j *jobMock) Loader() (string, interface{}) {
	return loaderFn()
}

func (j *jobMock) Handler(data []byte) error {
	return nil
}

func resetMocks() {
	enqueueFn = enqueueFnDefaultValue
	loaderFn = loaderFnDefaultValue
}

func TestDispatch(t *testing.T) {
	c := NewAsynqClientAdapter()
	c.client = &clientMock{}

	var want error

	got := c.Dispatch(&jobMock{})

	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

// Test for coverage purpose only
func TestDispatchWhenClientInstanceIsNil(t *testing.T) {
	c := NewAsynqClientAdapter()
	c.client = nil

	c.Dispatch(&jobMock{})
}

func TestFailToDispatchWhenInputCannotBeEncoded(t *testing.T) {
	defer t.Cleanup(resetMocks)

	c := NewAsynqClientAdapter()
	c.client = &clientMock{}

	loaderFn = func() (string, interface{}) {
		return "foo", make(chan int)
	}
	job := &jobMock{}

	got := c.Dispatch(job)

	_, want := got.(*json.UnsupportedTypeError)

	if !want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestFailToDispatchWhenClientCannotEnqueue(t *testing.T) {
	defer t.Cleanup(resetMocks)

	enqueueFn = func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
		return nil, fmt.Errorf("task cannot be nil")
	}

	c := NewAsynqClientAdapter()
	c.client = &clientMock{}

	got := c.Dispatch(&jobMock{})

	want := fmt.Errorf("task cannot be nil")

	if got.Error() != want.Error() {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
