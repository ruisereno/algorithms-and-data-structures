package queue

import (
	"errors"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()

	// Enqueue elements
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Validate size
	if q.Size() != 3 {
		t.Errorf("expected queue size 3, got %d", q.Size())
	}

	// Validate head and tail values
	head, _ := q.Peek()
	if head != 10 {
		t.Errorf("expected head value 10, got %d", head)
	}

	if q.tail.value != 30 {
		t.Errorf("expected tail value 30, got %d", q.tail.value)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue[int]()

	// Dequeue from empty queue
	_, err := q.Dequeue()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected 'queue is empty' error, got %v", err)
	}

	// Enqueue elements
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Dequeue elements and validate
	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("expected dequeued value 10, got %d", val)
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("expected dequeued value 20, got %d", val)
	}

	// Validate the last dequeue
	val, err = q.Dequeue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 30 {
		t.Errorf("expected dequeued value 30, got %d", val)
	}

	// Validate queue is empty after all dequeues
	if !q.IsEmpty() {
		t.Error("expected queue to be empty after dequeuing all elements")
	}
}

func TestQueuePeek(t *testing.T) {
	q := NewQueue[int]()

	// Peek on empty queue
	_, err := q.Peek()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected 'queue is empty' error, got %v", err)
	}

	// Enqueue elements
	q.Enqueue(10)
	q.Enqueue(20)

	// Validate Peek does not remove the element
	val, err := q.Peek()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("expected peek value 10, got %d", val)
	}

	// Ensure size remains the same after Peek
	if q.Size() != 2 {
		t.Errorf("expected queue size 2 after peek, got %d", q.Size())
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue[int]()

	// Initially, the queue should be empty
	if !q.IsEmpty() {
		t.Error("expected queue to be empty initially")
	}

	// Enqueue an element
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("expected queue to be non-empty after enqueue")
	}

	// Dequeue the element
	_, _ = q.Dequeue()
	if !q.IsEmpty() {
		t.Error("expected queue to be empty after dequeueing all elements")
	}
}

func TestQueueSize(t *testing.T) {
	q := NewQueue[int]()

	// Size of empty queue
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}

	// Enqueue elements
	q.Enqueue(10)
	q.Enqueue(20)
	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	// Dequeue an element
	_, _ = q.Dequeue()
	if q.Size() != 1 {
		t.Errorf("expected size 1 after one dequeue, got %d", q.Size())
	}

	// Dequeue last element
	_, _ = q.Dequeue()
	if q.Size() != 0 {
		t.Errorf("expected size 0 after dequeuing all elements, got %d", q.Size())
	}
}
