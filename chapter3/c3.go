package chapter3

// 3.2 Stack Min: How would you design a stack which, in addition to push and pop, has a function min

type LLNode[T int] struct {
	value T
	next  *LLNode[T]
}

type LLNodeWrapper[T int, K any] struct {
	value *LLNode[T]
	next  *LLNodeWrapper[T, K]
}

type Stack[T int] struct {
	stack *LLNode[T]
	min   *LLNodeWrapper[T, *LLNode[T]]
}

func (s *Stack[T]) Min() T {
	if s.min == nil {
		panic("Stack is empty")
	}
	return s.min.value.value
}

func NewStack[T int]() *Stack[T] {
	return &Stack[T]{
		stack: nil,
		min:   nil,
	}
}

func (s *Stack[T]) Push(value T) {
	newNode := &LLNode[T]{
		value: value,
		next:  s.stack,
	}
	s.stack = newNode

	if s.min == nil || value < s.Min() {
		s.min = &LLNodeWrapper[T, *LLNode[T]]{
			value: newNode,
			next:  s.min,
		}
	}
}

func (s *Stack[T]) Pop() T {
	if s.stack == nil {
		panic("Stack is empty")
	}
	res := s.stack
	s.stack = s.stack.next
	if res == s.min.value {
		s.min = s.min.next
	}
	return res.value
}

// 3.1 Three in One: Describe how you could use a single array to implement three stacks.
type TripleStack[T any] struct {
	array            []T
	lenA, lenB, lenC int
}

func NewTripleStack[T any]() *TripleStack[T] {
	return &TripleStack[T]{
		array: []T{},
	}
}

func (t *TripleStack[T]) PushA(value T) {
	newSlice := []T{value}
	newSlice = append(newSlice, t.array...)
	t.array = newSlice
	t.lenA++
}

func (t *TripleStack[T]) PushB(value T) {
	newSlice := make([]T, t.lenA)
	copy(newSlice, t.array[:t.lenA])
	newSlice = append(newSlice, value)
	newSlice = append(newSlice, t.array[t.lenA:]...)
	t.array = newSlice
	t.lenB++
}

func (t *TripleStack[T]) PushC(value T) {
	t.array = append(t.array, value)
	t.lenC++
}

func (t *TripleStack[T]) PopA() T {
	if t.lenA == 0 {
		panic("Stack A is empty")
	}
	value := t.array[0]
	newArr := t.array[1:]
	t.array = newArr
	t.lenA--
	return value
}

func (t *TripleStack[T]) PopB() T {
	if t.lenB == 0 {
		panic("Stack B is empty")
	}
	newArr := make([]T, t.lenA)
	copy(newArr, t.array[:t.lenA])
	value := t.array[t.lenA]
	newArr = append(newArr, t.array[t.lenA+1:]...)
	t.array = newArr
	t.lenB--
	return value
}

func (t *TripleStack[T]) PopC() T {
	if t.lenC == 0 {
		panic("Stack C is empty")
	}
	value := t.array[len(t.array)-1]
	newArr := t.array[:len(t.array)-1]
	t.array = newArr
	t.lenC--
	return value
}
