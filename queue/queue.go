// A -> B -> C
// A is head
// C is tail
// FIFO - Enqueue in tail and Dequeue in head
// Enqueue - From: A -> B -> C To: A -> B -> C -> D
// Dequeue - From: A -> B -> C To: B -> C

package queue

import (
	"errors"
)

var ErrEmpty = errors.New("queue is empty")

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{value: value}
	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, ErrEmpty
	}
	value := q.head.value
	q.head = q.head.next
	if q.IsEmpty() {
		q.tail = nil
	}
	q.size--
	return value, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, ErrEmpty
	}
	return q.head.value, nil
}

func (s *Queue[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Queue[T]) Size() int {
	return s.size
}

func (s *Queue[T]) LinearSeach(value T, comparator func(a, b T) bool) bool {
	current := s.head
	for current != nil {
		if comparator(current.value, value) {
			return true
		}
		current = current.next
	}
	return false
}
