package redirect

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/dorianneto/url-shortener/src/controller/redirect/input"
	"github.com/dorianneto/url-shortener/src/job"
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/gin-gonic/gin"
)

type QueueMock struct{}

func (q *QueueMock) Dispatch(job job.JobInterface) error {
	return nil
}

var createFn func(redirect *model.Redirect) (*model.Redirect, error) = func(redirect *model.Redirect) (*model.Redirect, error) {
	return redirect, nil
}

var findFn func(query input.FindRedirect) (*model.Redirect, error) = func(query input.FindRedirect) (*model.Redirect, error) {
	return &model.Redirect{Url: "https://example.com"}, nil
}

type RepositoryMock struct{}

func (rm *RepositoryMock) Find(query input.FindRedirect) (*model.Redirect, error) {
	return findFn(query)
}

func (rm *RepositoryMock) Create(redirect *model.Redirect) (*model.Redirect, error) {
	return createFn(redirect)
}

func TestRedirect(t *testing.T) {
	controller := &RedirectController{QueueClient: &QueueMock{}, Repository: &RepositoryMock{}}

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = httptest.NewRequest("GET", "/foo", new(bytes.Buffer))
	context.Params = append(context.Params, gin.Param{Key: "code", Value: "000000"})

	controller.Redirect(context)

	gotStatus := context.Writer.Status()
	wantStatus := 301

	if gotStatus != wantStatus {
		t.Errorf("Expected '%d', but got '%d'", wantStatus, gotStatus)
	}

	gotLocation := context.Writer.Header().Get("Location")
	wantLocation := "https://example.com"

	if gotLocation != wantLocation {
		t.Errorf("Expected '%s', but got '%s'", wantLocation, gotLocation)
	}
}
