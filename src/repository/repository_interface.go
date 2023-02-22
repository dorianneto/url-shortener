package repository

type RepositoryInterface interface {
	Find(code string) (interface{}, error)
	Create(data interface{}) (interface{}, error)
}
