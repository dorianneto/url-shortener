package redirect

import (
	"errors"
	"testing"

	"github.com/dorianneto/url-shortener/src/controller/redirect/input"
	"github.com/dorianneto/url-shortener/src/database/output/document"
	"github.com/dorianneto/url-shortener/src/model"
)

var readFn func() (*document.ReadOutput, error) = func() (*document.ReadOutput, error) {
	return &document.ReadOutput{Data: map[string]interface{}{
		"Url": "http://example.com",
	}}, nil
}

var writeFn func() (interface{}, error) = func() (interface{}, error) {
	return &model.Redirect{Code: "999999"}, nil
}

type DatabaseMock struct{}

func (d *DatabaseMock) Read(documentRef string) (*document.ReadOutput, error) {
	return readFn()
}

func (d *DatabaseMock) Write(documentRef string, data interface{}) (interface{}, error) {
	return writeFn()
}

func TestFind(t *testing.T) {
	r := RedirectRepository{
		Database: &DatabaseMock{},
	}

	input := input.FindRedirect{Code: "xxxxxx"}

	got, _ := r.Find(input)
	want := model.Redirect{Url: "http://example.com"}

	if got.Url != want.Url {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestFindWhenDataIsNotFound(t *testing.T) {
	readFn = func() (*document.ReadOutput, error) {
		return nil, errors.New("something goes wrong")
	}

	r := RedirectRepository{
		Database: &DatabaseMock{},
	}

	input := input.FindRedirect{Code: "xxxxxx"}

	_, got := r.Find(input)
	want := errors.New("something goes wrong")

	if got.Error() != want.Error() {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestCreate(t *testing.T) {
	r := RedirectRepository{
		Database: &DatabaseMock{},
	}

	got, _ := r.Create(&model.Redirect{Code: "999999"})
	want := &model.Redirect{Code: "999999"}

	if got.Code != want.Code {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestCreateWhenCannotSaveIntoDatabase(t *testing.T) {
	writeFn = func() (interface{}, error) {
		return nil, errors.New("something goes wrong")
	}

	r := RedirectRepository{
		Database: &DatabaseMock{},
	}

	_, got := r.Create(&model.Redirect{Code: "999999"})
	want := errors.New("something goes wrong")

	if got.Error() != want.Error() {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}
