package validator

import (
	"testing"
)

type sample struct {
	Foo string `validate:"required"`
	Bar int    `validate:"gt=10"`
}

func TestStructValidation(t *testing.T) {
	v := &ValidatorAdapter{}

	got := v.Struct(&sample{
		Foo: "lorem ipsum",
		Bar: 12,
	})

	var want error

	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}
