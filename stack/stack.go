// X -> Y -> Z
// X is head
// LIFO - Push in head and Pop from head
// Push - From: X -> Y -> Z To: W -> X -> Y -> Z
// Pop - From: X -> Y -> Z To: Y -> Z

package stack

import (
	"errors"
)

var ErrEmpty = errors.New("stack is empty")

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	head *Node[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{value: value}
	newNode.next = s.head
	s.head = newNode
	s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrEmpty
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrEmpty
	}
	return s.head.value, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) Size() int {
	return s.size
}
