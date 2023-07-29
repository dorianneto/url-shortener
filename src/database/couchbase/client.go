package couchbase

import (
	"log"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/dorianneto/url-shortener/src/database/output/document"
)

const DEFAULT_SCOPE = "redirects"
const DEFAULT_COLLECTION = "personal"

type GetResultAdapterInterface interface {
	Content(valuePtr interface{}) error
}

type CollectionAdapterInterface interface {
	Get(id string, opts *gocb.GetOptions) (docOut GetResultAdapterInterface, errOut error)
	Upsert(id string, val interface{}, opts *gocb.UpsertOptions) (mutOut *gocb.MutationResult, errOut error)
}

type collectionAdapter struct {
	collection *gocb.Collection
}

func (ca *collectionAdapter) Get(id string, opts *gocb.GetOptions) (docOut GetResultAdapterInterface, errOut error) {
	return ca.collection.Get(id, opts)
}

func (ca *collectionAdapter) Upsert(id string, val interface{}, opts *gocb.UpsertOptions) (mutOut *gocb.MutationResult, errOut error) {
	return ca.collection.Upsert(id, val, opts)
}

type ScopeAdapterInterface interface {
	Collection(collectionName string) CollectionAdapterInterface
}

type scopeAdapter struct {
	scope *gocb.Scope
}

func (sa *scopeAdapter) Collection(collectionName string) CollectionAdapterInterface {
	return &collectionAdapter{
		collection: sa.scope.Collection(collectionName),
	}
}

type BucketAdapterInterface interface {
	WaitUntilReady(timeout time.Duration, opts *gocb.WaitUntilReadyOptions) error
	Scope(scopeName string) ScopeAdapterInterface
}

type bucketAdater struct {
	bucket *gocb.Bucket
}

func (ba *bucketAdater) WaitUntilReady(timeout time.Duration, opts *gocb.WaitUntilReadyOptions) error {
	return ba.bucket.WaitUntilReady(timeout, opts)
}

func (ba *bucketAdater) Scope(scopeName string) ScopeAdapterInterface {
	return &scopeAdapter{
		scope: ba.bucket.Scope(scopeName),
	}
}

type CouchbaseClientAdapterInterface interface {
	Bucket(bucketName string) BucketAdapterInterface
	Close(opts *gocb.ClusterCloseOptions) error
}

type couchbaseClientAdapter struct {
	client *gocb.Cluster
}

func (cca *couchbaseClientAdapter) Bucket(bucketName string) BucketAdapterInterface {
	return &bucketAdater{
		bucket: cca.client.Bucket(bucketName),
	}
}

func (cca *couchbaseClientAdapter) Close(opts *gocb.ClusterCloseOptions) error {
	return cca.client.Close(opts)
}

type couchbaseAdapter struct {
	client     CouchbaseClientAdapterInterface
	bucketName string
}

func NewCouchbaseAdapter() *couchbaseAdapter {
	return &couchbaseAdapter{
		bucketName: os.Getenv("COUCHBASE_BUCKET"),
	}
}

func (ca *couchbaseAdapter) getClient() CouchbaseClientAdapterInterface {
	if ca.client != nil {
		return ca.client
	}

	connectionString := os.Getenv("COUCHBASE_HOST")
	username := os.Getenv("COUCHBASE_USERNAME")
	password := os.Getenv("COUCHBASE_PASSWORD")

	cluster, err := gocb.Connect("couchbase://"+connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	ca.client = &couchbaseClientAdapter{
		client: cluster,
	}

	return ca.client
}

func (ca *couchbaseAdapter) Close() error {
	return ca.getClient().Close(nil)
}

func (ca *couchbaseAdapter) Read(documentRef string) (*document.ReadOutput, error) {
	bucket := ca.getClient().Bucket(ca.bucketName)

	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, err
	}

	collection := bucket.Scope(DEFAULT_SCOPE).Collection(DEFAULT_COLLECTION)

	result, err := collection.Get(documentRef, nil)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = result.Content(&data)
	if err != nil {
		return nil, err
	}

	return &document.ReadOutput{Data: data}, nil
}

func (ca *couchbaseAdapter) Write(documentRef string, data interface{}) (interface{}, error) {
	bucket := ca.getClient().Bucket(ca.bucketName)

	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, err
	}

	collection := bucket.Scope(DEFAULT_SCOPE).Collection(DEFAULT_COLLECTION)

	_, err = collection.Upsert(documentRef, data, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}
