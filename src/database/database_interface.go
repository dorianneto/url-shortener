package database

type DatabaseInterface interface {
	Read(code string) (interface{}, error)
	Write(code string, data interface{}) (interface{}, error)
}
