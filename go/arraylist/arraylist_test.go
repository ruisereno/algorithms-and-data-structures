package arraylist

import (
	"errors"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	// Test with valid initial capacity
	list := NewArrayList[int](10)
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}
	if list.Capacity() != 10 {
		t.Errorf("expected capacity 10, got %d", list.Capacity())
	}

	// Test with invalid initial capacity (negative)
	list = NewArrayList[int](-1)
	if list.Capacity() != 5 {
		t.Errorf("expected default capacity 5 for negative initial capacity, got %d", list.Capacity())
	}

	// Test with invalid initial capacity (zero)
	list = NewArrayList[int](0)
	if list.Capacity() != 5 {
		t.Errorf("expected default capacity 5 for negative initial capacity, got %d", list.Capacity())
	}
}

func TestPush(t *testing.T) {
	list := NewArrayList[int](2)

	// Add elements
	list.Push(10)
	list.Push(20)

	// Verify size and capacity
	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}
	if list.Capacity() != 2 {
		t.Errorf("expected capacity 2, got %d", list.Capacity())
	}

	// Push beyond initial capacity
	list.Push(30)
	if list.Size() != 3 {
		t.Errorf("expected size 3, got %d", list.Size())
	}
	if list.Capacity() != 4 { // Capacity should double
		t.Errorf("expected capacity 4 after resize, got %d", list.Capacity())
	}
}

func TestPop(t *testing.T) {
	list := NewArrayList[int](5)

	// Pop an element from empty list
	_, err := list.Pop()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected 'array list is empty' error, got %v", err)
	}

	// Push elements and pop
	list.Push(10)
	list.Push(20)
	val, err := list.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("expected value 20, got %d", val)
	}
	if list.Size() != 1 {
		t.Errorf("expected size 1 after pop, got %d", list.Size())
	}
}

func TestGet(t *testing.T) {
	list := NewArrayList[int](5)

	// Push elements
	list.Push(10)
	list.Push(20)

	// Valid Get
	val, err := list.Get(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("expected value 20, got %d", val)
	}

	// Invalid Get (out of bounds)
	_, err = list.Get(2)
	if !errors.Is(err, ErrIndexOutOfBounds) {
		t.Errorf("expected 'index out of bounds' error, got %v", err)
	}
}

func TestSet(t *testing.T) {
	list := NewArrayList[int](5)

	// Push elements
	list.Push(10)
	list.Push(20)

	// Valid Set
	err := list.Set(1, 25)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	val, _ := list.Get(1)
	if val != 25 {
		t.Errorf("expected value 25, got %d", val)
	}

	// Invalid Set (out of bounds)
	err = list.Set(3, 30)
	if !errors.Is(err, ErrIndexOutOfBounds) {
		t.Errorf("expected 'index out of bounds' error, got %v", err)
	}
}

func TestSizeAndIsEmpty(t *testing.T) {
	list := NewArrayList[int](5)

	// Initially empty
	if !list.IsEmpty() {
		t.Error("expected list to be empty")
	}
	if list.Size() != 0 {
		t.Errorf("expected size 0, got %d", list.Size())
	}

	// Add elements
	list.Push(10)
	list.Push(20)
	if list.IsEmpty() {
		t.Error("expected list to be non-empty")
	}
	if list.Size() != 2 {
		t.Errorf("expected size 2, got %d", list.Size())
	}

	// Remove all elements
	list.Pop()
	list.Pop()
	if !list.IsEmpty() {
		t.Error("expected list to be empty after removing all elements")
	}
}

func TestLinearSeach(t *testing.T) {
	list := NewArrayList[int](0)

	// Linear search on an empty array list
	if list.LinearSeach(10, func(a, b int) bool {
		return a == b
	}) != false {
		t.Error("expected linear search to return false, got true")
	}

	// Add elements
	list.Push(10)
	list.Push(20)

	// Linear search for a value that is on the array list
	if list.LinearSeach(10, func(a, b int) bool {
		return a == b
	}) != true {
		t.Error("expected linear search to return true, got false")
	}

	// Linear search for a value that is not on the array list
	if list.LinearSeach(30, func(a, b int) bool {
		return a == b
	}) != false {
		t.Error("expected linear search to return false, got true")
	}
}
