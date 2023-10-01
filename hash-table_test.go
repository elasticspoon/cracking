package ctci

import (
	"fmt"
	"testing"
)

type HashNode struct {
	key   Hashable
	value any
}

func TestHashTableIntKeys(t *testing.T) {
	tests := []struct {
		insertions []HashNode
		deletions  []Hashable
		expected   []HashNode
		unexpected []HashNode
	}{
		{[]HashNode{{HashableInt(1), 1}}, []Hashable{HashableInt(1)}, []HashNode{}, []HashNode{{HashableInt(1), 1}}},
		{[]HashNode{{HashableString("test"), 1}}, []Hashable{HashableString("test")}, []HashNode{}, []HashNode{{HashableString("test"), 1}}},
	}

	for _, tt := range tests {

		hashTable := NewHashTable[Hashable, any](1)
		for _, v := range tt.insertions {
			hashTable.Set(v.key, v.value)
		}

		fmt.Println(hashTable.String())

		for _, v := range tt.insertions {
			if val, err := hashTable.Get(v.key); err != nil {
				t.Errorf("%v", err)
			} else if val != v.value {
				t.Errorf("Expected %v, got %v", v.value, val)
			}
		}

		for _, v := range tt.deletions {
			hashTable.Delete(v)
		}
		for _, v := range tt.expected {
			if val, err := hashTable.Get(v.key); err != nil {
				t.Errorf("%v", err)
			} else if val != v.value {
				t.Errorf("Expected %v, got %v", v.value, val)
			}
		}
		for _, v := range tt.unexpected {
			if v, err := hashTable.Get(v.key); err == nil {
				t.Errorf("Did not expect to find %v", v)
			}
		}
	}
}
