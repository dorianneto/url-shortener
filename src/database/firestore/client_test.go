package firestore

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"cloud.google.com/go/firestore"
)

var (
	haveCloseBeenCalled = false
	getFn, setFn        func() error
)

type documentAdapterMock struct{}

func (a *documentAdapterMock) Get(ctx context.Context) (_ *firestore.DocumentSnapshot, err error) {
	return &firestore.DocumentSnapshot{}, getFn()
}

func (a *documentAdapterMock) Set(ctx context.Context, data interface{}, opts ...firestore.SetOption) (_ *firestore.WriteResult, err error) {
	return nil, setFn()
}

type clientAdapterMock struct{}

func (c *clientAdapterMock) Doc(path string) documentAdapterInterface {
	return &documentAdapterMock{}
}

func (c *clientAdapterMock) Close() error {
	haveCloseBeenCalled = true
	return nil
}

func resetMocks() {
	haveCloseBeenCalled = false
	getFn = func() error {
		return nil
	}
	setFn = func() error {
		return nil
	}
}

func TestClientClose(t *testing.T) {
	defer t.Cleanup(resetMocks)

	c := NewFirestoreAdapter()
	c.client = &clientAdapterMock{}

	c.Close()

	if !haveCloseBeenCalled {
		t.Error("Method Close() is expected to be called")
	}
}

func TestClientRead(t *testing.T) {
	defer t.Cleanup(resetMocks)

	c := NewFirestoreAdapter()
	c.client = &clientAdapterMock{}

	got, err := c.Read("foo")

	if err != nil {
		t.Errorf("An error is not expected")
	}

	if reflect.TypeOf(got).String() != "*document.ReadOutput" {
		t.Error("The return expected to be instance of *document.ReadOutput")
	}
}

func TestClientCannotReadBecauseBecauseSomethingGoesWrong(t *testing.T) {
	defer t.Cleanup(resetMocks)

	getFn = func() error {
		return errors.New("something goes wrong")
	}

	c := NewFirestoreAdapter()
	c.client = &clientAdapterMock{}

	got, err := c.Read("foo")

	if err == nil {
		t.Errorf("An error is expected")
	}

	if got != nil {
		t.Error("The return expected to be nil")
	}
}

func TestClientWrite(t *testing.T) {
	defer t.Cleanup(resetMocks)

	c := NewFirestoreAdapter()
	c.client = &clientAdapterMock{}

	got, err := c.Write("foo", map[string]int{"foo": 1, "bar": 2})

	if err != nil {
		t.Errorf("An error is not expected")
	}

	_, ok := got.(map[string]int)

	if !ok {
		t.Error("The return expected to be map[string]int")
	}
}

func TestClientCannotWriteBecauseSomethingGoesWrong(t *testing.T) {
	defer t.Cleanup(resetMocks)

	setFn = func() error {
		return errors.New("something goes wrong")
	}

	c := NewFirestoreAdapter()
	c.client = &clientAdapterMock{}

	_, err := c.Write("foo", nil)

	if err == nil {
		t.Errorf("An error is expected")
	}
}
