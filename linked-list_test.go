package ctci

import "testing"

func TestLinkedList(t *testing.T) {
	tests := []struct {
		insertions  []int
		deletions   []int
		expected    []int
		failsToFind []int
	}{
		{[]int{1}, []int{}, []int{1}, []int{}},
		{[]int{1}, []int{1}, []int{}, []int{}},
		{[]int{1, 2, 3, 4}, []int{1}, []int{2, 3, 4}, []int{}},
		{[]int{1, 2, 3, 4}, []int{5, 7, 5}, []int{1, 2, 3, 4}, []int{}},
		{[]int{1, 2, 3, 4}, []int{}, []int{1, 2, 3, 4}, []int{5, 7, 9}},
	}

	for _, tt := range tests {
		list := LinkedList[int, int]{nil, 0}
		for _, v := range tt.insertions {
			list.Insert(v)
		}

		for _, v := range tt.deletions {
			list.Delete(v)
		}

		for _, v := range tt.expected {
			list.Search(v)
			if _, err := list.Search(v); err != nil {
				t.Errorf("%v", err)
			}
		}
		for _, v := range tt.failsToFind {
			list.Search(v)
			if _, err := list.Search(v); err == nil {
				t.Errorf("%v", err)
			}
		}
	}
}
