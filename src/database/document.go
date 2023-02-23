package database

import "github.com/dorianneto/url-shortener/src/database/output/document"

type DocumentInterface interface {
	Read(documentRef string) (*document.ReadOutput, error)
	Write(documentRef string, data interface{}) (interface{}, error)
}
