package ctci

type ArrayList[T any] struct {
	array    []T
	Length   int
	Capacity int
}

func (al *ArrayList[T]) Append(el T) ArrayList[T] {
	if al.Length == al.Capacity {
		al.Capacity *= 2
		temp := make([]T, al.Capacity)
		copy(temp, al.array)
		al.array = temp
	}
	al.array[al.Length] = el
	al.Length++
	return ArrayList[T]{al.array, al.Length, al.Capacity}
}

func (al *ArrayList[T]) At(i int) T {
	return al.array[i]
}

func NewArrayList[T any]() ArrayList[T] {
	return ArrayList[T]{make([]T, 1), 0, 1}
}
