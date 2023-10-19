package chapter3

import "testing"

func TestCalculator(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"1+2", 3},
		{"1-2", -1},
		{"1*2", 2},
		{"1/2", 0.5},
		{"1+2*3", 7},
		{"1+2*3/2", 4},
		{"1+2*3/2-5", -1},
		{"1+2*3/2-5*2", -6},
	}

	c := NewCalculator()
	for _, test := range tests {
		result := c.Calculate(test.input)
		if result != test.expected {
			t.Errorf("Expected %s = %f, got %f", test.input, test.expected, result)
		}
	}
}
