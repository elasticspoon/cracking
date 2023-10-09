package chapter3

import (
	"reflect"
	"testing"
)

func TestTripleStack(t *testing.T) {
	tests := []struct {
		action        string
		value         int
		expectedState []int
	}{
		{"pushC", 1, []int{1}},
		{"pushA", 4, []int{4, 1}},
		{"pushA", 3, []int{3, 4, 1}},
		{"pushC", 3, []int{3, 4, 1, 3}},
		{"pushB", 0, []int{3, 4, 0, 1, 3}},
		{"pushB", 7, []int{3, 4, 7, 0, 1, 3}},
		{"popC", 3, []int{3, 4, 7, 0, 1}},
		{"popA", 3, []int{4, 7, 0, 1}},
		{"popB", 7, []int{4, 0, 1}},
		{"popB", 0, []int{4, 1}},
		{"popA", 4, []int{1}},
		{"popC", 1, []int{}},
		{"pushA", 1, []int{1}},
		{"pushB", 3, []int{1, 3}},
		{"pushA", 2, []int{2, 1, 3}},
		{"pushC", 5, []int{2, 1, 3, 5}},
		{"pushB", 4, []int{2, 1, 4, 3, 5}},
		{"pushA", 6, []int{6, 2, 1, 4, 3, 5}},
		{"popC", 5, []int{6, 2, 1, 4, 3}},
		{"popB", 4, []int{6, 2, 1, 3}},
		{"popA", 6, []int{2, 1, 3}},
		{"popA", 2, []int{1, 3}},
		{"popB", 3, []int{1}},
		{"popA", 1, []int{}},
	}

	stack := NewTripleStack[int]()
	for _, test := range tests {
		value := test.value
		switch test.action {
		case "pushA":
			stack.PushA(test.value)
		case "pushB":
			stack.PushB(test.value)
		case "pushC":
			stack.PushC(test.value)
		case "popA":
			value = stack.PopA()
		case "popB":
			value = stack.PopB()
		case "popC":
			value = stack.PopC()
		}
		if !reflect.DeepEqual(stack.array, test.expectedState) {
			t.Errorf("Expected %v, got %v", test.expectedState, stack.array)
		}
		if value != test.value {
			t.Errorf("Expected %v, got %v", test.value, value)
		}
	}
}

func TestMinStack(t *testing.T) {
	tests := []struct {
		action string
		value  int
	}{
		{"push", 8},
		{"min", 8},
		{"push", 5},
		{"min", 5},
		{"push", 7},
		{"min", 5},
		{"push", 8},
		{"push", 9},
		{"pop", 9},
		{"min", 5},
		{"pop", 8},
		{"pop", 7},
		{"min", 5},
		{"pop", 5},
		{"min", 8},
		{"pop", 8},
	}

	stack := NewStack[int]()
	for _, test := range tests {
		value := test.value
		switch test.action {
		case "push":
			stack.Push(test.value)
		case "pop":
			value = stack.Pop()
		case "min":
			value = stack.Min()
		}

		if value != test.value {
			t.Errorf("Expected %v, got %v", test.value, value)
		}
	}
}
