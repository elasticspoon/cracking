package builtins

import (
	"fmt"
	"math"
)

// find min/max in a heap is O(1): peek
// push is O(log n): insert at the end and bubble up
// pop is O(log n): replace root with last element and bubble down
//
// NewHeap() -> Heap
// heapify([]int) -> Heap
//
// swap(Node, Node) bool: swap two nodes, return true if swapped, false otherwise
// siftUp(Node) bool: sift up the node, return true if sifted, false otherwise
// siftDown(Node) bool: sift down the node, return true if sifted, false otherwise
type Heap struct {
	array []BasicNode
}

type BasicNode interface {
	Key() int
	Value() int
}

func NewHeap() *Heap {
	return &Heap{array: []BasicNode{}}
}

func (h *Heap) IsEmpty() bool {
	return len(h.array) == 0
}

func (h *Heap) Len() int {
	return len(h.array)
}

func (h *Heap) Heapify(input []BasicNode) {
	for _, v := range input {
		h.Push(v)
	}
}

func (h *Heap) Peek() any {
	return h.array[0].Value()
}

func (h *Heap) Push(v BasicNode) {
	h.array = append(h.array, v)
	h.siftUp(len(h.array) - 1)
}

func (h *Heap) String() string {
	return fmt.Sprintf("%v", h.array)
}

func (h *Heap) Pop() any {
	val := h.array[0]
	h.array[0], h.array[len(h.array)-1] = h.array[len(h.array)-1], h.array[0]
	h.array = h.array[:len(h.array)-1]

	h.siftDown(0)
	return (val).Value()
}

func (h *Heap) siftUp(i int) {
	for child, parent := i, (i-1)/2; parent >= 0; {
		if h.array[child].Value() < h.array[parent].Value() {
			h.array[child], h.array[parent] = h.array[parent], h.array[child]
			child = parent
			parent = (parent - 1) / 2
		} else {
			break
		}
	}
}

func (h *Heap) siftDown(parent int) {
	if len(h.array) <= parent {
		return
	}

	for {
		childA := 2*parent + 1
		childB := 2*parent + 2

		var valA, valB int
		valP := h.array[parent].Value()

		if childA >= len(h.array) {
			return
		} else if childB >= len(h.array) {
			valA = h.array[childA].Value()
			valB = math.MaxInt
		} else {
			valA, valB = h.array[childA].Value(), h.array[childB].Value()
		}

		if valP < valA && valP < valB {
			return
		}

		if valA < valB {
			h.array[childA], h.array[parent] = h.array[parent], h.array[childA]
			parent = childA
			// fmt.Println("broke here? 1")
		} else {
			fmt.Println(h.array)
			h.array[childB], h.array[parent] = h.array[parent], h.array[childB]
			parent = childB
			// fmt.Println("broke here? 2")
		}
	}
}
