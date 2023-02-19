package database

type DatabaseInterface interface {
	Read() (interface{}, error)
	Write() (interface{}, error)
}
