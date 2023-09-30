package ctci

type HashTable[K comparable, V comparable] struct {
	list []LinkedList[K, V]
	Size int
}

func (ht *HashTable[K, V]) Set(key K, value V) {
}

func (ht *HashTable[K, V]) Delete(el K) {
}

func NewHashTable[K comparable, V comparable]() HashTable[K, V] {
	return HashTable[K, V]{make([]LinkedList[K, V], 0), 0}
}

func (ht *HashTable[K, V]) Get(key K) (*V, error) {
	linkedList := ht.list[ht.hash(key)]

	if node, err := linkedList.Search(key); err != nil {
		return nil, err
	} else {
		return &node.value, nil
	}
}

func (ht *HashTable[K, V]) hash(key K) int {
	return 0
}
