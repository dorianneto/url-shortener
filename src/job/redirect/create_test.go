package redirect

import (
	"errors"
	"testing"

	"github.com/dorianneto/url-shortener/src/controller/redirect/input"
	"github.com/dorianneto/url-shortener/src/model"
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

	job := &CreateRedirectJob{
		Payload: r,
	}
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
	data := []byte(`{"id":"9961c85e-d805-4a96-bbdb-7ab35ef67ca8","url":"https://example.com","code":"PkevciZPWVHk","created_at":"0001-01-01T00:00:00Z"}`)

	job := &CreateRedirectJob{
		Repository: &RedirectRepositoryMock{},
	}
	got := job.Handler(data)

	var want error

	if got != want {
		t.Errorf("Expected '%v', but got '%s'", want, got)
	}
}

func TestHandlerWhenPayloadIsNotWellFormed(t *testing.T) {
	data := []byte(`{"id":10,"url":"https://example.com","code":"PkevciZPWVHk","created_at":"0001-01-01T00:00:00Z"}`)

	job := &CreateRedirectJob{
		Repository: &RedirectRepositoryMock{},
	}
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

	data := []byte(`{"id":"9961c85e-d805-4a96-bbdb-7ab35ef67ca8","url":"https://example.com","code":"PkevciZPWVHk","created_at":"0001-01-01T00:00:00Z"}`)

	job := &CreateRedirectJob{
		Repository: &RedirectRepositoryMock{},
	}
	got := job.Handler(data)

	var want error

	if got == want {
		t.Errorf("Expected '%v', but got '%s'", want, got)
	}
}
