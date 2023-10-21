package builtins

import (
	"testing"
)

type HeapNode struct {
	key   int
	value int
}

func (n *HeapNode) Key() int   { return n.key }
func (n *HeapNode) Value() int { return n.value }

func TestHeapify(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		h := NewHeap()
		h.Heapify(toHeapNodes(tt.input))
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
		h.Heapify(toHeapNodes(tt.input))

		result := []int{}

		for !h.IsEmpty() {
			result = append(result, h.Pop().(int))
		}

		for i, v := range tt.expected {
			if result[i] != v {
				t.Errorf("Pop() want %v, got %v", tt.expected, result)
				break
			}
		}
	}
}

func isHeap(a []BasicNode) bool {
	for i := 0; i < len(a); i++ {
		if i*2+1 < len(a) && a[i].Value() > a[i*2+1].Value() {
			return false
		}
		if i*2+2 < len(a) && a[i].Value() > a[i*2+2].Value() {
			return false
		}
	}
	return true
}

func toHeapNodes(in []int) []BasicNode {
	var out []BasicNode
	for _, v := range in {
		out = append(out, &HeapNode{key: v, value: v})
	}
	return out
}
