package database

type DocumentInterface interface {
	Read(documentRef string) (*ReadDocumentOutput, error)
	Write(documentRef string, data interface{}) (interface{}, error)
}
