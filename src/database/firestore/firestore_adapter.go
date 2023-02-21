package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
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

func (fa *FilestoreAdapter) Read() (interface{}, error) {
	data := fa.getClient().Doc("States/California")

	document, err := data.Get(fa.contextBackground)
	if err != nil {
		return nil, err
	}

	return document.Data(), nil
}

func (fa *FilestoreAdapter) Write(code string, data interface{}) (interface{}, error) {
	document := fa.getClient().Doc("Redirects/" + code)

	_, err := document.Set(fa.contextBackground, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
