package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/dorianneto/url-shortener/internal/database/output/document"
)

type documentAdapterInterface interface {
	Get(ctx context.Context) (_ *firestore.DocumentSnapshot, err error)
	Set(ctx context.Context, data interface{}, opts ...firestore.SetOption) (_ *firestore.WriteResult, err error)
}

type FirestoreClientAdapterInterface interface {
	Doc(path string) documentAdapterInterface
	Close() error
}

type firestoreClientAdapter struct {
	client *firestore.Client
}

func (fca *firestoreClientAdapter) Close() error {
	return fca.client.Close()
}

func (fca *firestoreClientAdapter) Doc(path string) documentAdapterInterface {
	return fca.client.Doc(path)
}

type firestoreAdapter struct {
	client            FirestoreClientAdapterInterface
	contextBackground context.Context
}

func NewFirestoreAdapter() *firestoreAdapter {
	return &firestoreAdapter{}
}

func (fa *firestoreAdapter) getClient() FirestoreClientAdapterInterface {
	if fa.client != nil {
		return fa.client
	}

	fa.contextBackground = context.Background()
	client, err := firestore.NewClient(fa.contextBackground, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		log.Fatalln(err)
	}

	fa.client = &firestoreClientAdapter{
		client: client,
	}

	return fa.client
}

func (fa *firestoreAdapter) Close() error {
	return fa.getClient().Close()
}

func (fa *firestoreAdapter) Read(documentRef string) (*document.ReadOutput, error) {
	data := fa.getClient().Doc("Redirects/" + documentRef)

	result, err := data.Get(fa.contextBackground)
	if err != nil {
		return nil, err
	}

	return &document.ReadOutput{Data: result.Data()}, nil
}

func (fa *firestoreAdapter) Write(documentRef string, data interface{}) (interface{}, error) {
	document := fa.getClient().Doc("Redirects/" + documentRef)

	_, err := document.Set(fa.contextBackground, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
