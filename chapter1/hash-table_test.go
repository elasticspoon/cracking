package ctci

import (
	"testing"
)

type HashNode struct {
	key   Hashable
	value any
}

func TestHashTableRehash(t *testing.T) {
	tests := []struct {
		insertions int
		deletions  int
		sizeMax    int
		sizeMin    int
	}{
		{1, 0, 2, 2},
		{10, 0, 16, 16},
	}

	for _, tt := range tests {

		hashTable := NewHashTable[Hashable, any](1)
		for i := 0; i < tt.insertions; i++ {
			hashTable.Set(HashableInt(i), i)
		}

		if hashTable.Size != tt.sizeMax {
			t.Errorf("Expected size %v, got %v", tt.sizeMax, hashTable.Size)
		}

		for i := 0; i < tt.deletions; i++ {
			hashTable.Delete(HashableInt(i))
		}

		if hashTable.Size != tt.sizeMin {
			t.Errorf("Expected size %v, got %v", tt.sizeMin, hashTable.Size)
		}
	}
}

func TestHashTableItems(t *testing.T) {
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
