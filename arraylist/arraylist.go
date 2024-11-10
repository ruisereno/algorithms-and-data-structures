package arraylist

import "errors"

var (
	ErrEmpty            = errors.New("array list is empty")
	ErrIndexOutOfBounds = errors.New("index out of bounds")
)

type ArrayList[T any] struct {
	data []T
	size int
}

func NewArrayList[T any](initialCapacity int) *ArrayList[T] {
	if initialCapacity <= 0 {
		initialCapacity = 5
	}
	return &ArrayList[T]{
		data: make([]T, initialCapacity),
		size: 0,
	}
}

func (a *ArrayList[T]) resize() {
	newData := make([]T, len(a.data)*2)
	copy(newData, a.data)
	a.data = newData
}

func (a *ArrayList[T]) Push(value T) {
	if a.size == len(a.data) {
		a.resize()
	}
	a.data[a.size] = value
	a.size++
}

func (a *ArrayList[T]) Pop() (T, error) {
	if a.IsEmpty() {
		var zero T
		return zero, ErrEmpty
	}
	value := a.data[a.size-1]
	a.size--
	return value, nil
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= a.size {
		var zero T
		return zero, ErrIndexOutOfBounds
	}
	return a.data[index], nil
}

func (a *ArrayList[T]) Set(index int, value T) error {
	if index < 0 || index >= a.size {
		return ErrIndexOutOfBounds
	}
	a.data[index] = value
	return nil
}

func (a *ArrayList[T]) Size() int {
	return a.size
}

func (a *ArrayList[T]) Capacity() int {
	return len(a.data)
}

func (a *ArrayList[T]) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayList[T]) LinearSeach(value T, comparator func(a, b T) bool) bool {
	for i := 0; i < a.size; i++ {
		if comparator(a.data[i], value) {
			return true
		}
	}
	return false
}
