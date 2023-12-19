package chapter8

import "testing"

func TestNumJumps(t *testing.T) {
	tests := []struct {
		steps int
		want  int
	}{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 7},
		{5, 13},
		{6, 24},
		{7, 44},
	}

	for _, test := range tests {
		got := numJumps(test.steps)
		if got != test.want {
			t.Errorf("NumJumps(%d) = %d", test.steps, got)
		}
	}
}

func TestMagicIndexDups(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0}, 0},
		{[]int{5, 5, 5, 5, 5, 5, 5, 5, 5}, 5},
		{[]int{-20, -11, 0, 3, 4, 5, 6, 8, 9, 11}, 4},
		{[]int{-20, -11, 2, 3, 4, 5, 6, 8, 9, 11}, 4},
		{[]int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}, 2},
		{[]int{-10, -5, 1, 2, 2, 3, 4, 7, 9, 12, 13}, 7},
		{[]int{9, 9, 9, 9}, -1},
		{[]int{6, 6, 6, 6, 6, 6, 6}, 6},
	}
	for _, test := range tests {
		got := magicIndexDups(test.input)
		if got != test.want {
			t.Errorf("MagicIndexDups(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestMagicIndex(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{-20, -11, 0, 3, 4, 5, 6, 8, 9, 11}, 4},
		{[]int{3, 4, 5, 6, 8, 9, 11}, -1},
		{[]int{0}, 0},
		{[]int{0, 1, 2, 3, 4, 5}, 2},
	}
	for _, test := range tests {
		got := magicIndex(test.input)
		if got != test.want {
			t.Errorf("MagicIndex(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestMultiplyWithoutSign(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
		{20, 30, 600},
		{20, 5, 100},
		{139923023, 9672943, 139923023 * 9672943},
	}

	for _, test := range tests {
		got := multiplyWithoutSign(test.a, test.b)
		if got != test.want {
			t.Errorf("MultiplyWithoutSign(%d, %d) = %d, want %d", test.a, test.b, got, test.want)
		}
	}
}
