package stack

import (
	"errors"
	"testing"
)

func TestStackPush(t *testing.T) {
	s := NewStack[int]()

	// Push elements
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("expected stack size 3, got %d", s.Size())
	}

	// Check top element after push
	top, _ := s.Peek()
	if top != 3 {
		t.Errorf("expected top element 3, got %d", top)
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack[int]()

	// Pop from empty stack
	_, err := s.Pop()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected 'stack is empty' error, got %v", err)
	}

	// Push elements
	s.Push(1)
	s.Push(2)

	// Pop elements and validate
	val, err := s.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("expected popped value 2, got %d", val)
	}

	val, err = s.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("expected popped value 1, got %d", val)
	}

	// Pop again to ensure error
	_, err = s.Pop()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected 'stack is empty' error, got %v", err)
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack[int]()

	// Peek on empty stack
	_, err := s.Peek()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected 'stack is empty' error, got %v", err)
	}

	// Push and Peek
	s.Push(42)
	val, err := s.Peek()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 42 {
		t.Errorf("expected peek value 42, got %d", val)
	}

	// Ensure Peek does not remove the element
	if s.Size() != 1 {
		t.Errorf("expected size 1 after peek, got %d", s.Size())
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := NewStack[int]()

	// Check empty stack
	if !s.IsEmpty() {
		t.Error("expected stack to be empty")
	}

	// Push an element
	s.Push(1)
	if s.IsEmpty() {
		t.Error("expected stack to be non-empty")
	}

	// Pop the element
	_, _ = s.Pop()
	if !s.IsEmpty() {
		t.Error("expected stack to be empty after popping last element")
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack[int]()

	// Size of empty queue
	if s.Size() != 0 {
		t.Errorf("expected initial size 0, got %d", s.Size())
	}

	// Push elements
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("expected size 2 after two pushes, got %d", s.Size())
	}

	// Pop an element
	_, _ = s.Pop()
	if s.Size() != 1 {
		t.Errorf("expected size 1 after one pop, got %d", s.Size())
	}

	// Pop last element
	_, _ = s.Pop()
	if s.Size() != 0 {
		t.Errorf("expected size 0 after popping all elements, got %d", s.Size())
	}
}

func TestLinearSeach(t *testing.T) {
	s := NewStack[int]()

	// Linear search on an empty stack
	if s.LinearSeach(10, func(a, b int) bool {
		return a == b
	}) != false {
		t.Error("expected linear search to return false, got true")
	}

	// Push elements
	s.Push(10)
	s.Push(20)

	// Linear search for a value that is on the stack
	if s.LinearSeach(10, func(a, b int) bool {
		return a == b
	}) != true {
		t.Error("expected linear search to return true, got false")
	}

	// Linear search for a value that is not on the stack
	if s.LinearSeach(30, func(a, b int) bool {
		return a == b
	}) != false {
		t.Error("expected linear search to return false, got true")
	}
}
