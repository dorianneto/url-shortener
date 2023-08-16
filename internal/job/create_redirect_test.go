package job

import (
	"errors"
	"testing"

	"github.com/dorianneto/url-shortener/internal/controller/redirect/input"
	"github.com/dorianneto/url-shortener/internal/model"
)

var createFn func(redirect *model.Redirect) (*model.Redirect, error) = func(redirect *model.Redirect) (*model.Redirect, error) {
	return redirect, nil
}

type RedirectRepositoryMock struct{}

func (rrm *RedirectRepositoryMock) Find(query input.FindRedirect) (*model.Redirect, error) {
	return nil, nil
}

func (rrm *RedirectRepositoryMock) Create(redirect *model.Redirect) (*model.Redirect, error) {
	return createFn(redirect)
}

func TestLoader(t *testing.T) {
	r, _ := model.NewRedirect("https://example.com")

	job := NewCreateRedirectJob(&RedirectRepositoryMock{})
	job.LoadPayload(r)
	gotQueueName, gotPayload := job.Loader()

	want := []interface{}{
		"create:redirect",
		r,
	}

	if gotQueueName != want[0] {
		t.Errorf("Expected '%s', but got '%s'", want, gotQueueName)
	}

	if gotPayload != want[1] {
		t.Errorf("Expected '%v', but got '%v'", want, gotPayload)
	}
}

func TestHandler(t *testing.T) {
	r, _ := model.NewRedirect("https://example.com")
	data := []byte(`{"id":"9961c85e-d805-4a96-bbdb-7ab35ef67ca8","url":"https://example.com","code":"PkevciZPWVHk","created_at":"0001-01-01T00:00:00Z"}`)

	job := NewCreateRedirectJob(&RedirectRepositoryMock{})
	job.LoadPayload(r)
	got := job.Handler(data)

	var want error

	if got != want {
		t.Errorf("Expected '%v', but got '%s'", want, got)
	}
}

func TestHandlerWhenPayloadIsNotWellFormed(t *testing.T) {
	r, _ := model.NewRedirect("https://example.com")
	data := []byte(`{"id":10,"url":"https://example.com","code":"PkevciZPWVHk","created_at":"0001-01-01T00:00:00Z"}`)

	job := NewCreateRedirectJob(&RedirectRepositoryMock{})
	job.LoadPayload(r)
	got := job.Handler(data)

	var want error

	if got == want {
		t.Errorf("Expected '%v', but got '%s'", want, got)
	}
}

func TestHandlerWhenRepositoryGoesWrong(t *testing.T) {
	createFn = func(redirect *model.Redirect) (*model.Redirect, error) {
		return nil, errors.New("")
	}

	r, _ := model.NewRedirect("https://example.com")
	data := []byte(`{"id":"9961c85e-d805-4a96-bbdb-7ab35ef67ca8","url":"https://example.com","code":"PkevciZPWVHk","created_at":"0001-01-01T00:00:00Z"}`)

	job := NewCreateRedirectJob(&RedirectRepositoryMock{})
	job.LoadPayload(r)
	got := job.Handler(data)

	var want error

	if got == want {
		t.Errorf("Expected '%v', but got '%s'", want, got)
	}
}
