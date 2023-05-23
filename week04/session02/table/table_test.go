package table

import "testing"

func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 1, 2, "+", 3, ""},
		{"subtraction", 1, 2, "-", -1, ""},
		{"multiplication", 1, 2, "*", 2, ""},
		{"division", 1, 2, "/", 0, ""},
		{"division by zero", 1, 0, "/", 0, "cannot divide by zero"},
		{"invalid operator", 1, 2, "x", 0, "invalid operator"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			actual, err := DoMath(d.num1, d.num2, d.op)
			if actual != d.expected {
				t.Errorf("expected %d, got %d", d.expected, actual)
			}
			if err != nil && err.Error() != d.errMsg {
				t.Errorf("expected %s, got %s", d.errMsg, err.Error())
			}
		})
	}
}
