package repository

type RepositoryInterface interface {
	Find() (interface{}, error)
	Create() (interface{}, error)
}
