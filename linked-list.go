package ctci

import "fmt"

func (ll *LinkedListError[T]) Error() string {
	return fmt.Sprintf("key %v not found", ll.key)
}

type LinkedListError[K comparable] struct {
	key K
}

type node[K comparable, V any] struct {
	key   K
	value V
	next  *node[K, V]
}

func (n *node[K, V]) delete(key K) *node[K, V] {
	if n.key == key {
		return n.next
	} else if n.next == nil {
		return n
	} else {
		n.next = n.next.delete(key)
		return n
	}
}

func (n *node[K, V]) search(key K) (*node[K, V], error) {
	if n.key == key {
		return n, nil
	} else if n.next == nil {
		return nil, &LinkedListError[K]{key}
	} else {
		return n.next.search(key)
	}
}

type LinkedList[K comparable, V any] struct {
	head   *node[K, V]
	Length int
}

func (ll *LinkedList[K, V]) Delete(key K) {
	if ll.head != nil {
		ll.Length--
		ll.head = ll.head.delete(key)
	}
}

func (ll *LinkedList[K, V]) Insert(key K, values ...V) {
	ll.Length++
	if len(values) > 0 {
		ll.head = &node[K, V]{key: key, next: ll.head, value: values[0]}
	} else {
		ll.head = &node[K, V]{key: key, next: ll.head}
	}
}

func (ll *LinkedList[K, V]) Search(el K) (*node[K, V], error) {
	if ll.head == nil {
		return nil, &LinkedListError[K]{el}
	} else {
		return ll.head.search(el)
	}
}
