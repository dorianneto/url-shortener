package model

import (
	"reflect"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	r := new(Redirect)

	got := len(r.generateCode())
	want := 12

	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestConstructor(t *testing.T) {
	got, _ := NewRedirect("https://example.com")

	want := new(Redirect)

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestConstructorWhenUrlIsEmpty(t *testing.T) {
	got, _ := NewRedirect("")

	var want *Redirect

	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
