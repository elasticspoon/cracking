package ctci

import "fmt"

func (ll *LinkedListError[T]) Error() string {
	return fmt.Sprintf("value %v not found", ll.value)
}

type LinkedListError[T comparable] struct {
	value T
}

type node[T comparable] struct {
	value T
	next  *node[T]
}

func (n *node[T]) delete(el T) *node[T] {
	if n.value == el {
		return n.next
	} else if n.next == nil {
		return n
	} else {
		n.next = n.next.delete(el)
		return n
	}
}

func (n *node[T]) search(el T) (*node[T], error) {
	if n.value == el {
		return n, nil
	} else if n.next == nil {
		return nil, &LinkedListError[T]{el}
	} else {
		return n.next.search(el)
	}
}

type LinkedList[T comparable] struct {
	head   *node[T]
	Length int
}

func (ll *LinkedList[T]) Delete(el T) {
	if ll.head != nil {
		ll.head = ll.head.delete(el)
	}
}

func (ll *LinkedList[T]) Insert(el T) {
	ll.head = &node[T]{el, ll.head}
}

func (ll *LinkedList[T]) Search(el T) (*node[T], error) {
	if ll.head == nil {
		return nil, &LinkedListError[T]{el}
	} else {
		return ll.head.search(el)
	}
}
