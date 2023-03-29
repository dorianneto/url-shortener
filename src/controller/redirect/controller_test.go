package redirect

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dorianneto/url-shortener/src/controller/redirect/input"
	"github.com/dorianneto/url-shortener/src/job"
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/gin-gonic/gin"
)

const (
	URL_FAKE = "https://example.com"
)

type QueueMock struct{}

func (q *QueueMock) Dispatch(job job.JobInterface) error {
	return nil
}

var createFn func(redirect *model.Redirect) (*model.Redirect, error) = func(redirect *model.Redirect) (*model.Redirect, error) {
	return redirect, nil
}

var findFn func(query input.FindRedirect) (*model.Redirect, error) = func(query input.FindRedirect) (*model.Redirect, error) {
	return &model.Redirect{Url: URL_FAKE}, nil
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
	wantLocation := URL_FAKE

	if gotLocation != wantLocation {
		t.Errorf("Expected '%s', but got '%s'", wantLocation, gotLocation)
	}
}

func TestRedirectWhenInputIsEmpty(t *testing.T) {
	controller := &RedirectController{QueueClient: &QueueMock{}, Repository: &RepositoryMock{}}

	request := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(request)
	context.Request = httptest.NewRequest("GET", "/foo", new(bytes.Buffer))
	context.Params = append(context.Params, gin.Param{Key: "code", Value: ""})

	controller.Redirect(context)

	gotStatus := context.Writer.Status()
	wantStatus := 400

	if gotStatus != wantStatus {
		t.Errorf("Expected '%d', but got '%d'", wantStatus, gotStatus)
	}

	gotErrorMessage := request.Body.String()
	wantErrorMessage := `{"message":"Key: 'FindRedirect.Code' Error:Field validation for 'Code' failed on the 'required' tag"}`

	if gotErrorMessage != wantErrorMessage {
		t.Errorf("Expected '%s', but got '%s'", wantErrorMessage, gotErrorMessage)
	}
}

func TestRedirectWhenDataIsNotFoundInDatabase(t *testing.T) {
	findFn = func(query input.FindRedirect) (*model.Redirect, error) {
		return nil, errors.New("something goes wrong")
	}

	controller := &RedirectController{QueueClient: &QueueMock{}, Repository: &RepositoryMock{}}

	request := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(request)
	context.Request = httptest.NewRequest("GET", "/foo", new(bytes.Buffer))
	context.Params = append(context.Params, gin.Param{Key: "code", Value: "000000"})

	controller.Redirect(context)

	gotStatus := context.Writer.Status()
	wantStatus := 400

	if gotStatus != wantStatus {
		t.Errorf("Expected '%d', but got '%d'", wantStatus, gotStatus)
	}

	gotErrorMessage := request.Body.String()
	wantErrorMessage := `{"message":"something goes wrong"}`

	if gotErrorMessage != wantErrorMessage {
		t.Errorf("Expected '%s', but got '%s'", wantErrorMessage, gotErrorMessage)
	}
}

func TestStore(t *testing.T) {
	controller := &RedirectController{QueueClient: &QueueMock{}, Repository: &RepositoryMock{}}

	request := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(request)
	context.Request = httptest.NewRequest("POST", "/", strings.NewReader(`{"url":"`+URL_FAKE+`"}`))

	controller.Store(context)

	gotStatus := context.Writer.Status()
	wantStatus := 201

	if gotStatus != wantStatus {
		t.Errorf("Expected '%d', but got '%d'", wantStatus, gotStatus)
	}
}

func TestStoreWhenPayloadIsEmpty(t *testing.T) {
	controller := &RedirectController{QueueClient: &QueueMock{}, Repository: &RepositoryMock{}}

	request := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(request)
	context.Request = httptest.NewRequest("POST", "/", strings.NewReader("{}"))

	controller.Store(context)

	gotStatus := context.Writer.Status()
	wantStatus := 400

	if gotStatus != wantStatus {
		t.Errorf("Expected '%d', but got '%d'", wantStatus, gotStatus)
	}
}

func TestStoreWhenModelCannotBeCreated(t *testing.T) {
	controller := &RedirectController{QueueClient: &QueueMock{}, Repository: &RepositoryMock{}}

	request := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(request)
	context.Request = httptest.NewRequest("POST", "/", strings.NewReader(`{"url":"foo"}`))

	controller.Store(context)

	gotStatus := context.Writer.Status()
	wantStatus := 400

	if gotStatus != wantStatus {
		t.Errorf("Expected '%d', but got '%d'", wantStatus, gotStatus)
	}
}
