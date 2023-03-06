package database

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/dorianneto/url-shortener/src/database/output/document"
)

type FirestoreAdapter struct {
	client            *firestore.Client
	contextBackground context.Context
}

func (fa *FirestoreAdapter) getClient() *firestore.Client {
	if fa.client != nil {
		return fa.client
	}

	fa.contextBackground = context.Background()
	client, err := firestore.NewClient(fa.contextBackground, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		log.Fatalln(err)
	}

	fa.client = client

	return fa.client
}

func (fa *FirestoreAdapter) Close() error {
	return fa.getClient().Close()
}

func (fa *FirestoreAdapter) Read(documentRef string) (*document.ReadOutput, error) {
	data := fa.getClient().Doc("Redirects/" + documentRef)

	result, err := data.Get(fa.contextBackground)
	if err != nil {
		return nil, err
	}

	return &document.ReadOutput{Data: result.Data()}, nil
}

func (fa *FirestoreAdapter) Write(documentRef string, data interface{}) (interface{}, error) {
	document := fa.getClient().Doc("Redirects/" + documentRef)

	_, err := document.Set(fa.contextBackground, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
