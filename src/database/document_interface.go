package database

type DocumentInterface interface {
	Read(documentRef string) (interface{}, error)
	Write(documentRef string, data interface{}) (interface{}, error)
}
