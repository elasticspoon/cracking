package chapter5

import "testing"

func TestBookNextLarger(t *testing.T) {
	tests := []struct {
		given    int
		expected int
	}{
		{0b11101, 0b11110},
		{0b11011010101011, 0b11011010101101},
		{0b100, 0b1000},
		{0b1111000, 0b10000111},
		{0b111, 0b1011},
		{0b0, -1},
	}

	for _, tt := range tests {
		actual := bookNextLarger(tt.given)
		if actual != tt.expected {
			t.Errorf("nextLarger(%b): expected %b, actual %b", tt.given, tt.expected, actual)
		}
	}
}

func TestNextLarger(t *testing.T) {
	tests := []struct {
		given    int
		expected int
	}{
		{0b11101, 0b11110},
		{0b11011010101011, 0b11011010101101},
		{0b100, 0b1000},
		{0b1111000, 0b10000111},
		{0b111, 0b1011},
		{0b0, -1},
	}

	for _, tt := range tests {
		actual := nextLarger(tt.given)
		if actual != tt.expected {
			t.Errorf("nextLarger(%b): expected %b, actual %b", tt.given, tt.expected, actual)
		}
	}
}

func TestNextSmaller(t *testing.T) {
	tests := []struct {
		given    int
		expected int
	}{
		{0b0, -1},
		{0b1, -1},
		{0b111, -1},
		{0b1011, 0b111},
		{0b1000, 0b100},
		{0b10000111, 0b1000111},
		{0b10110111, 0b10101111},
	}

	for _, tt := range tests {
		actual := nextSmaller(tt.given)
		if actual != tt.expected {
			t.Errorf("nextLarger(%b): expected %b, actual %b", tt.given, tt.expected, actual)
		}
	}
}

func TestBitsToFlip(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{0b11101, 0b11110, 2},
		{0b11011010101011, 0b11011010101101, 2},
		{0b100, 0b1000, 2},
		{0b1111000, 0b10000111, 8},
		{0b10000111, 0b1111000, 8},
		{0b111, 0b1011, 2},
		{0b1011, 0b111, 2},
	}

	for _, tt := range tests {
		actual := BitsToFlip(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("bitsToFlip(%b, %b): expected %d, actual %d", tt.a, tt.b, tt.expected, actual)
		}
	}
}

func TestSwapOddEvenBits(t *testing.T) {
	tests := []struct {
		given    int
		expected int
	}{
		{0b10101010, 0b01010101},
		{0b11100011, 0b11010011},
		{0b0, 0b0},
		{0b1, 0b10},
		{0b11, 0b11},
		{0b111, 0b1011},
		{0b1111, 0b1111},
		{-1, -1},
	}

	for _, tt := range tests {
		actual := SwapOddEvenBits(tt.given)
		if actual != tt.expected {
			t.Errorf("swapOddEvenBits(%b): expected %b, actual %b", tt.given, tt.expected, actual)
		}
	}
}
