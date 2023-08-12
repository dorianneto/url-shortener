package couchbase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/couchbase/gocb/v2"
)

var (
	haveCloseBeenCalled = false
	readyFn             func() error
	scopeFn             func() ScopeAdapterInterface
	getFn               func() (docOut GetResultAdapterInterface, errOut error)
	upsertFn            func() (mutOut *gocb.MutationResult, errOut error)
	contentFn           func() error
)

type contentAdapterMock struct{}

func (cam *contentAdapterMock) Content(valuePtr interface{}) error {
	return contentFn()
}

type collectionAdapterMock struct{}

func (cam *collectionAdapterMock) Get(id string, opts *gocb.GetOptions) (docOut GetResultAdapterInterface, errOut error) {
	return getFn()
}

func (cam *collectionAdapterMock) Upsert(id string, val interface{}, opts *gocb.UpsertOptions) (mutOut *gocb.MutationResult, errOut error) {
	return upsertFn()
}

type scopeAdapterMock struct{}

func (sam *scopeAdapterMock) Collection(collectionName string) CollectionAdapterInterface {
	return &collectionAdapterMock{}
}

type bucketAdapterMock struct{}

func (a *bucketAdapterMock) WaitUntilReady(timeout time.Duration, opts *gocb.WaitUntilReadyOptions) error {
	return readyFn()
}

func (a *bucketAdapterMock) Scope(scopeName string) ScopeAdapterInterface {
	return scopeFn()
}

type clientAdapterMock struct{}

func (c *clientAdapterMock) Bucket(bucketName string) BucketAdapterInterface {
	return &bucketAdapterMock{}
}

func (c *clientAdapterMock) Close(opts *gocb.ClusterCloseOptions) error {
	haveCloseBeenCalled = true
	return nil
}

func resetMocks() {
	haveCloseBeenCalled = false
	contentFn = func() error {
		return nil
	}
	getFn = func() (docOut GetResultAdapterInterface, errOut error) {
		return &contentAdapterMock{}, nil
	}
	upsertFn = func() (mutOut *gocb.MutationResult, errOut error) {
		return nil, nil
	}
	scopeFn = func() ScopeAdapterInterface {
		return &scopeAdapterMock{}
	}
	readyFn = func() error {
		return nil
	}
}

func TestClientClose(t *testing.T) {
	defer t.Cleanup(resetMocks)

	c := NewCouchbaseAdapter()
	c.client = &clientAdapterMock{}

	c.Close()

	if !haveCloseBeenCalled {
		t.Error("Method Close() is expected to be called")
	}
}

func TestClientRead(t *testing.T) {
	defer t.Cleanup(resetMocks)

	c := NewCouchbaseAdapter()
	c.client = &clientAdapterMock{}

	got, err := c.Read("foo")

	if err != nil {
		t.Errorf("An error is not expected")
	}

	if reflect.TypeOf(got).String() != "*document.ReadOutput" {
		t.Error("The return expected to be instance of *document.ReadOutput")
	}
}

func TestClientCannotReadBecauseItIsNotReady(t *testing.T) {
	defer t.Cleanup(resetMocks)

	readyFn = func() error {
		return errors.New("something goes wrong")
	}

	c := NewCouchbaseAdapter()
	c.client = &clientAdapterMock{}

	got, err := c.Read("foo")

	if err == nil {
		t.Errorf("An error is expected")
	}

	if got != nil {
		t.Error("The return expected to be nil")
	}
}

func TestClientCannotReadWhenDocumentCannotBeFound(t *testing.T) {
	defer t.Cleanup(resetMocks)

	getFn = func() (docOut GetResultAdapterInterface, errOut error) {
		return nil, errors.New("something goes wrong")
	}

	c := NewCouchbaseAdapter()
	c.client = &clientAdapterMock{}

	got, err := c.Read("foo")

	if err == nil {
		t.Errorf("An error is expected")
	}

	if got != nil {
		t.Error("The return expected to be nil")
	}
}

func TestClientCannotReadWhenContentCannotBeAssigned(t *testing.T) {
	defer t.Cleanup(resetMocks)

	contentFn = func() error {
		return errors.New("something goes wrong")
	}

	c := NewCouchbaseAdapter()
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

	c := NewCouchbaseAdapter()
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

func TestClientCannotWriteBecauseItIsNotReady(t *testing.T) {
	defer t.Cleanup(resetMocks)

	readyFn = func() error {
		return errors.New("something goes wrong")
	}

	c := NewCouchbaseAdapter()
	c.client = &clientAdapterMock{}

	_, err := c.Write("foo", nil)

	if err == nil {
		t.Errorf("An error is expected")
	}
}

func TestClientCannotWriteBecauseInsertWentWrongOnDatabaseSide(t *testing.T) {
	defer t.Cleanup(resetMocks)

	upsertFn = func() (mutOut *gocb.MutationResult, errOut error) {
		return nil, errors.New("something goes wrong")
	}

	c := NewCouchbaseAdapter()
	c.client = &clientAdapterMock{}

	_, err := c.Write("foo", nil)

	if err == nil {
		t.Errorf("An error is expected")
	}
}
