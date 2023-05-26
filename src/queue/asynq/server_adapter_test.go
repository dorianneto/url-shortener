package asynq

import (
	"context"
	"errors"
	"testing"

	"github.com/dorianneto/url-shortener/src/job/jobtest"
	"github.com/hibiken/asynq"
)

var (
	serverRunError   error
	isShutdownCalled = false
)

type serverMock struct{}

func (sm *serverMock) Run(handler asynq.Handler) error {
	return serverRunError
}

func (sm *serverMock) Shutdown() {
	isShutdownCalled = true
}

type serverMuxMock struct{}

func (sm *serverMuxMock) HandleFunc(pattern string, handler func(context.Context, *asynq.Task) error) {
}

func (sm *serverMuxMock) ProcessTask(context.Context, *asynq.Task) error {
	return nil
}

func resetServerMocks() {
	isShutdownCalled = false
	serverRunError = nil
}

func TestRegisterWorker(t *testing.T) {
	defer t.Cleanup(resetServerMocks)

	s := NewAsynqServerdapter()

	s.RegisterWorker(&jobtest.JobMock{})

	got := s.workers

	_, isJobMock := got[0].(*jobtest.JobMock)

	if len(got) <= 0 {
		t.Error("Expected workers to have more than 0")
	}

	if !isJobMock {
		t.Error("Expected worker to be an instance from jobtest.JobMock")
	}
}

func TestRunWorkers(t *testing.T) {
	defer t.Cleanup(resetServerMocks)

	s := NewAsynqServerdapter()
	s.server = &serverMock{}
	s.mux = &serverMuxMock{}

	s.RegisterWorker(&jobtest.JobMock{})
	s.RunWorkers()

	if isShutdownCalled {
		t.Error("Shutdown method from server object SHOULD NOT been called")
	}
}

func TestRunWorkersFailWhenRunningServerGoesWrong(t *testing.T) {
	defer t.Cleanup(resetServerMocks)

	serverRunError = errors.New("something goes wrong")

	s := NewAsynqServerdapter()
	s.server = &serverMock{}
	s.mux = &serverMuxMock{}

	s.RegisterWorker(&jobtest.JobMock{})
	s.RunWorkers()

	if !isShutdownCalled {
		t.Error("Shutdown method from server object SHOULD been called")
	}
}

func TestRunWorkersFailWhenNotPassingRedisConcurrencyValue(t *testing.T) {
	defer t.Cleanup(resetServerMocks)

	s := NewAsynqServerdapter()
	s.server = nil

	s.RunWorkers()

	if s.mux != nil {
		t.Error("Expected that error when getting server instance stop the executation")
	}
}
