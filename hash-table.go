package ctci

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
)

type (
	HashableInt    int
	HashableString string
)

type Hashable interface {
	Hash() []byte
	Equal(Hashable) bool
}

func (hi HashableInt) Hash() []byte {
	arr := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(arr, int64(hi))
	return arr[:n]
}

func (hi HashableInt) Equal(h Hashable) bool {
	v, ok := h.(HashableInt)
	return ok && v == hi
}

func (hs HashableString) Hash() []byte {
	return []byte(hs)
}

func (hs HashableString) Equal(h Hashable) bool {
	v, ok := h.(HashableString)
	return ok && v == hs
}

type HashLLError[K Hashable] struct {
	key K
}

func (ll *HashLLError[K]) Error() string {
	return fmt.Sprintf("value %v not found", ll.key)
}

type hashLLnode[K Hashable] struct {
	value any
	key   K
	next  *hashLLnode[K]
}

func (n *hashLLnode[T]) delete(key T) *hashLLnode[T] {
	if n.key.Equal(key) {
		return n.next
	} else if n.next == nil {
		return n
	} else {
		n.next = n.next.delete(key)
		return n
	}
}

func (n *hashLLnode[K]) search(key K) (*hashLLnode[K], error) {
	if n.key.Equal(key) {
		return n, nil
	} else if n.next == nil {
		return nil, &HashLLError[K]{key}
	} else {
		return n.next.search(key)
	}
}

type HashLL[T Hashable] struct {
	head   *hashLLnode[T]
	Length int
}

func (ll *HashLL[T]) Delete(key T) {
	if ll.head != nil {
		ll.head = ll.head.delete(key)
	}
}

func (ll *HashLL[T]) Insert(key T, value any) {
	ll.head = &hashLLnode[T]{key: key, value: value, next: ll.head}
}

func (ll *HashLL[T]) Search(key T) (*hashLLnode[T], error) {
	if ll.head == nil {
		return nil, &HashLLError[T]{key}
	} else {
		return ll.head.search(key)
	}
}

type HashTable[K Hashable, V any] struct {
	list  []HashLL[K]
	Items int
	Size  int
}

func (ht *HashTable[K, V]) Set(key K, value any) {
	bucket := &ht.list[ht.hash(key)]

	bucket.Insert(key, value)
	ht.Items++
	ht.rehash()
}

func (ht *HashTable[K, V]) Delete(key K) {
	bucket := &ht.list[ht.hash(key)]
	bucket.Delete(key)
	ht.Items--
	ht.rehash()
}

func NewHashTable[K Hashable, V any](size int) HashTable[K, V] {
	if size < 1 {
		panic("size must be greater than 0")
	}

	return HashTable[K, V]{make([]HashLL[K], size), 0, size}
}

func (ht *HashTable[K, V]) Get(key K) (any, error) {
	linkedList := ht.list[ht.hash(key)]

	if node, err := linkedList.Search(key); err != nil {
		return nil, err
	} else {
		return node.value, nil
	}
}

func (ht *HashTable[K, V]) hash(key K) int {
	hasher := sha256.New()
	hasher.Write(key.Hash())
	hashed := hasher.Sum(nil)
	hashedInt := new(big.Int).SetBytes(hashed)
	return int(hashedInt.Int64()) % ht.Size
}

func (ht *HashTable[K, V]) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	for _, v := range ht.list {
		for walker := v.head; walker != nil; walker = walker.next {
			out.WriteString(fmt.Sprintf("%v -> %v,", walker.key, walker.value))
		}
	}
	out.WriteString("}")
	return out.String()
}

func (ht *HashTable[K, V]) rehash() {
	var newSize int

	if density := float64(ht.Items) / float64(ht.Size); density > 0.8 {
		newSize = ht.Size * 2
	} else if density < 0.2 {
		newSize = ht.Size / 2
	} else {
		return
	}

	newHashTable := NewHashTable[K, V](newSize)

	for _, bucket := range ht.list {
		walker := bucket.head
		for walker != nil {
			newHashTable.Set(walker.key, walker.value)
			walker = walker.next
		}
	}

	ht.list = newHashTable.list
	ht.Size = newSize
	ht.Items = newHashTable.Items
}
