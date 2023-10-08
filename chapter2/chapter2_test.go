package ctci

import (
	"testing"
)

func TestAddListsFwd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{0, 0, 0},
		{312, 426, 837},
		{312, 926, 842},
		{111, 222, 333},
	}

	for _, tt := range tests {
		a := revList(intToRevList(tt.a))
		b := revList(intToRevList(tt.b))
		got := AddListsFwd(a, b)

		if got != tt.want {
			t.Errorf("AddListsFwd(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestAddListsRev(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{0, 0, 0},
		{312, 426, 738},
		{312, 926, 1238},
		{111, 222, 333},
	}

	for _, tt := range tests {
		a := intToRevList(tt.a)
		b := intToRevList(tt.b)
		got := AddListsRev(a, b)

		if got != tt.want {
			t.Errorf("AddListsRev(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestListIsPal(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"a", true},
		{"ab", false},
		{"aa", true},
		{"aba", true},
		{"abba", true},
		{"abcba", true},
		{"abcd", false},
	}

	for _, tt := range tests {
		list := stringToList(tt.s)
		got := ListIsPal(list)
		if got != tt.want {
			t.Errorf("ListIsPal(%s) = %t, want %t", tt.s, got, tt.want)
		}
	}
}

func TestFindCycle(t *testing.T) {
	tests := []struct {
		s    string
		want rune
	}{
		{"able", 'e'},
		{"abcd", 'd'},
		{"bcdekh", 'h'},
	}

	for _, tt := range tests {
		list := stringToList(tt.s)

		for curr := list; curr != nil; curr = curr.next {
			if curr.next == nil {
				curr.next = list
				break
			}
		}

		got := FindCycle(list)
		if got.value != tt.want {
			t.Errorf("FindCycle(%s) = %c, want %c", tt.s, got.value, tt.want)
		}
	}
}

func revList(head *LLNode[int]) *LLNode[int] {
	var newHead *LLNode[int]
	for ; head != nil; head = head.next {
		newHead = &LLNode[int]{value: head.value, next: newHead}
	}

	return newHead
}

func stringToList(s string) *LLNode[any] {
	var list *LLNode[any]
	for _, c := range s {
		list = &LLNode[any]{value: c, next: list}
	}
	return list
}

func intToRevList(v int) *LLNode[int] {
	var list *LLNode[int]
	for v > 0 {
		list = &LLNode[int]{value: v % 10, next: list}
		v /= 10
	}
	return list
}
