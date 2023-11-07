package chapter4

import "testing"

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		input    [][2]int
		expected []int
	}{
		{[][2]int{{1, 2}}, []int{1, 2}},
		{[][2]int{{1, 2}, {2, 3}}, []int{1, 2, 3}},
		{[][2]int{{2, 1}, {3, 2}, {4, 3}}, []int{4, 3, 2, 1}},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}}, []int{1, 2, 3, 4}},
		// TODO: fix this?
		// I dont know how to make this test work. Cause the sorts are not always the same.
		// But are valid
		// {[][2]int{{5, 11}, {7, 11}, {7, 8}, {3, 8}, {3, 10}, {11, 2}, {11, 9}, {11, 10}, {8, 9}}, []int{7, 3, 5, 8, 11, 2, 9, 10}},
	}

	for _, test := range tests {
		actual := TopologicalSort(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("TopologicalSort(%v) = %v, want %v", test.input, actual, test.expected)
		} else {
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("TopologicalSort(%v) = %v, want %v", test.input, actual, test.expected)
					break
				}
			}
		}
	}
}
