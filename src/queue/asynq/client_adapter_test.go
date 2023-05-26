package asynq

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dorianneto/url-shortener/src/job/jobtest"
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

func resetClientMocks() {
	enqueueFn = enqueueFnDefaultValue
	jobtest.LoaderFn = jobtest.LoaderFnDefaultValue
}

func TestDispatch(t *testing.T) {
	c := NewAsynqClientAdapter()
	c.client = &clientMock{}

	var want error

	got := c.Dispatch(&jobtest.JobMock{})

	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

// Test for coverage purpose only
func TestDispatchWhenClientInstanceIsNil(t *testing.T) {
	c := NewAsynqClientAdapter()
	c.client = nil

	c.Dispatch(&jobtest.JobMock{})
}

func TestFailToDispatchWhenInputCannotBeEncoded(t *testing.T) {
	defer t.Cleanup(resetClientMocks)

	c := NewAsynqClientAdapter()
	c.client = &clientMock{}

	jobtest.LoaderFn = func() (string, interface{}) {
		return "foo", make(chan int)
	}
	job := &jobtest.JobMock{}

	got := c.Dispatch(job)

	_, want := got.(*json.UnsupportedTypeError)

	if !want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestFailToDispatchWhenClientCannotEnqueue(t *testing.T) {
	defer t.Cleanup(resetClientMocks)

	enqueueFn = func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
		return nil, fmt.Errorf("task cannot be nil")
	}

	c := NewAsynqClientAdapter()
	c.client = &clientMock{}

	got := c.Dispatch(&jobtest.JobMock{})

	want := fmt.Errorf("task cannot be nil")

	if got.Error() != want.Error() {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
