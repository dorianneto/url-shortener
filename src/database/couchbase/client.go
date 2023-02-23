package database

import (
	"log"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/dorianneto/url-shortener/src/database/output/document"
)

const DEFAULT_SCOPE = "redirects"
const DEFAULT_COLLECTION = "personal"

type CouchbaseAdapter struct {
	client     *gocb.Cluster
	bucketName string
}

func (ca *CouchbaseAdapter) getClient() *gocb.Cluster {
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

	ca.bucketName = os.Getenv("COUCHBASE_BUCKET")
	ca.client = cluster

	return ca.client
}

func (ca *CouchbaseAdapter) Close() error {
	return ca.getClient().Close(nil)
}

func (ca *CouchbaseAdapter) Read(documentRef string) (*document.ReadOutput, error) {
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

func (ca *CouchbaseAdapter) Write(documentRef string, data interface{}) (interface{}, error) {
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
