package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/dorianneto/url-shortener/src/database"
)

type FilestoreAdapter struct {
	client            *firestore.Client
	contextBackground context.Context
}

func (fa *FilestoreAdapter) getClient() *firestore.Client {
	if fa.client != nil {
		return fa.client
	}

	fa.contextBackground = context.Background()
	client, err := firestore.NewClient(fa.contextBackground, "dumb-project-id")
	if err != nil {
		log.Fatalln(err)
	}

	fa.client = client

	return fa.client
}

func (fa *FilestoreAdapter) Close() error {
	return fa.getClient().Close()
}

func (fa *FilestoreAdapter) Read(documentRef string) (*database.ReadDocumentOutput, error) {
	data := fa.getClient().Doc("Redirects/" + documentRef)

	document, err := data.Get(fa.contextBackground)
	if err != nil {
		return nil, err
	}

	return &database.ReadDocumentOutput{Data: document.Data()}, nil
}

func (fa *FilestoreAdapter) Write(documentRef string, data interface{}) (interface{}, error) {
	document := fa.getClient().Doc("Redirects/" + documentRef)

	_, err := document.Set(fa.contextBackground, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
