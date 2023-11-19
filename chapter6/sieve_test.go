package chapter6

import "testing"

func TestIthPrime(t *testing.T) {
	tests := []struct {
		given    int
		expected int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 11},
		{6, 13},
		{7, 17},
		{8, 19},
		{9, 23},
		{10, 29},
		{30, 113},
		{100, 541},
		{1000, 7919},
		// {10000, 104729}, // should panic
	}

	for _, tt := range tests {
		if got := ithPrime(tt.given); got != tt.expected {
			t.Errorf("ithPrime(%v) = %v, want %v", tt.given, got, tt.expected)
		}
	}
}
