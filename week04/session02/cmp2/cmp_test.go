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

	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	result := CreatePerson("John", 42)
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Errorf(diff)
	}
}
