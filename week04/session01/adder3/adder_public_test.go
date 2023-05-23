package adder_test

import (
	"test_examples/adder3"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	if adder.AddNumbers(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}
