package ctci

import "testing"

func TestHashTable(t *testing.T) {
	tests := []struct {
		insertions [][2]int
		deletions  []int
		expected   [][2]int
		unexpected [][2]int
	}{
		{([][2]int{{1, 1}}), ([]int{1}), ([][2]int{{1, 1}}), ([][2]int{})},
		{([][2]int{{1, 1}}), ([]int{1}), ([][2]int{{1, 1}}), ([][2]int{})},
	}

	for _, tt := range tests {

		hashTable := NewHashTable[int, int]()
		for _, v := range tt.insertions {
			hashTable.Set(v[0], v[1])
		}

		for _, v := range tt.insertions {
			if val, err := hashTable.Get(v[0]); err != nil {
			} else if val != v[1] {
			}
		}

		for _, v := range tt.deletions {
			hashTable.Delete(v)
		}
	}
}
