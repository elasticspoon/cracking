package builtins

import (
	"testing"
)

func TestHeapify(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		h := NewHeap()
		h.Heapify(tt.input)
		got := h.array
		if !isHeap(got) {
			t.Errorf("Heapify(%v) got %v", tt.input, got)
		}
	}
}

func TestPopHeap(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 1, 4, 2, 3}, []int{1, 2, 3, 4, 5}},
		{[]int{9, 34, 11, 2, 3, 5, 1, 4, 2, 3}, []int{1, 2, 2, 3, 3, 4, 5, 9, 11, 34}},
	}

	for _, tt := range tests {
		h := NewHeap()
		h.Heapify(tt.input)

		result := []int{}

		for !h.IsEmpty() {
			result = append(result, h.Pop())
		}

		for i, v := range tt.expected {
			if result[i] != v {
				t.Errorf("Pop() want %v, got %v", tt.expected, result)
				break
			}
		}
	}
}

func isHeap(a []int) bool {
	for i := 0; i < len(a); i++ {
		if i*2+1 < len(a) && a[i] > a[i*2+1] {
			return false
		}
		if i*2+2 < len(a) && a[i] > a[i*2+2] {
			return false
		}
	}
	return true
}
