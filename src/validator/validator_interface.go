package validator

type ValidatorInterface interface {
	Struct(input interface{}) error
}
