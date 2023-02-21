package database

type DatabaseInterface interface {
	Read() (interface{}, error)
	Write(code string, data interface{}) (interface{}, error)
}
