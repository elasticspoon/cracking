package chapter5

import (
	"reflect"
	"testing"
)

func TestInsertBit(t *testing.T) {
	tests := []struct {
		n    uint32
		m    uint32
		i    int
		j    int
		want uint32
	}{
		{0b10000000000, 0b10011, 2, 6, 0b10001001100},
	}

	for _, tt := range tests {
		if got := InsertBitsBetween(tt.n, tt.m, tt.i, tt.j); got != tt.want {
			t.Errorf("InsertBit(%v, %v, %v, %v) = %v, want %v", tt.n, tt.m, tt.i, tt.j, got, tt.want)
		}
	}
}

func TestFloatDecToBinary(t *testing.T) {
	tests := []struct {
		n    float64
		want string
	}{
		{0.5, "0.1"},
		{0.75, "0.11"},
		{0.625, "0.101"},
		{0.1, "ERROR"},
	}

	for _, tt := range tests {
		if got := FloatDecToBinary(tt.n); got != tt.want {
			t.Errorf("FloatDecToBinary(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestFlipBit(t *testing.T) {
	tests := []struct {
		given int
		want  int
	}{
		{0b11011101111, 8},
		{0b1111, 4},
		{0b1110, 4},
	}

	for _, tt := range tests {
		if got := FlipBit(tt.given); got != tt.want {
			t.Errorf("FlipBit(%v) = %v, want %v", tt.given, got, tt.want)
		}
	}
}

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

func TestScreenWriter(t *testing.T) {
	t.Run("test screen", func(t *testing.T) {
		tests := []struct {
			given            []byte
			width, x1, x2, y int
			expectedScreen   []byte
		}{
			{[]byte{0, 0, 0}, 8, 1, 1, 2, []byte{0, 128, 0}},
			{[]byte{0, 0, 0}, 8, 2, 4, 2, []byte{0, 0b01110000, 0}},
			{[]byte{0, 0, 0}, 8, 1, 8, 2, []byte{0, 255, 0}},
			{[]byte{0, 0, 0}, 24, 1, 24, 1, []byte{255, 255, 255}},
			{[]byte{0, 0, 0}, 24, 2, 23, 1, []byte{0b01111111, 255, 0b11111110}},
		}

		for _, tt := range tests {
			actualScreen, actualErr := DrawLine(tt.given, tt.width, tt.x1, tt.x2, tt.y)
			if actualErr == false {
				t.Errorf("screenWriter(%v, %d, %d, %d, %d): expected no error", tt.given, tt.width, tt.x1, tt.x2, tt.y)
			}

			if reflect.DeepEqual(actualScreen, tt.expectedScreen) == false {
				t.Errorf("screenWriter(%v, %d, %d, %d, %d): expected %v, actual %v", tt.given, tt.width, tt.x1, tt.x2, tt.y, tt.expectedScreen, actualScreen)
			}
		}
	})
	t.Run("test errors", func(t *testing.T) {
		tests := []struct {
			given            []byte
			width, x1, x2, y int
			expectedScreen   []byte
			expectedErr      bool
		}{
			{[]byte{}, 8, 0, 0, 0, []byte{}, false},          // empty screen
			{[]byte{2, 3, 4}, 32, 0, 0, 0, []byte{}, false},  // row not long enough
			{[]byte{2, 3, 4}, 8, 0, 0, 4, []byte{}, false},   // y is out of bounds
			{[]byte{2, 3, 4}, 8, 1, 0, 2, []byte{}, false},   // x1 > x2
			{[]byte{2, 3, 4}, 8, 1, 0, 2, []byte{}, false},   // x1 > x2
			{[]byte{2, 3, 4}, 8, 11, 12, 2, []byte{}, false}, // x is out of bounds right
			{[]byte{2, 3, 4}, 8, 0, 0, 2, []byte{}, false},   //  x is out of bounds left
			{[]byte{2, 3, 4}, 8, -1, 1, 2, []byte{}, false},  //  x is out of bounds left
		}

		for _, tt := range tests {
			_, actualErr := DrawLine(tt.given, tt.width, tt.x1, tt.x2, tt.y)
			if actualErr != tt.expectedErr {
				t.Errorf("screenWriter(%v, %d, %d, %d, %d): expected %v, actual %v", tt.given, tt.width, tt.x1, tt.x2, tt.y, tt.expectedErr, actualErr)
			}
		}
	})
}
