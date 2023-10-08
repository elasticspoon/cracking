package chapter3

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
