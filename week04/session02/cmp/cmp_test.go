package cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "John",
		Age:  42,
	}
	result := CreatePerson("John", 42)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf(diff)
	}
}
