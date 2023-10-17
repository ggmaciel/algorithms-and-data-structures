package doublylinkedlist

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList[int]()

	_, err := dll.RemoveAt(1)

	if err == nil {
		t.Errorf("Expected out of bounds error, got %v", err)
	}

	dll.Append(5)
	dll.Append(7)
	dll.Append(9)

	getValue, _ := dll.Get(2)

	if *getValue != 9 {
		t.Errorf("Expected 9, got %d", *getValue)
	}

	removeValue, _ := dll.RemoveAt(1)

	if *removeValue != 7 {
		t.Errorf("Expected 7, got %d", *removeValue)
	}

	length := dll.Len()

	if length != 2 {
		t.Errorf("Expected 2, got %d", length)
	}

	dll.Append(11)
	removeValue, _ = dll.RemoveAt(1)

	if *removeValue != 9 {
		t.Errorf("Expected 9, got %d", *removeValue)
	}

	removeValue = dll.Remove(9)

	if removeValue != nil {
		t.Errorf("Expected nil, got %d", *removeValue)
	}

	removeValue, _ = dll.RemoveAt(0)

	if *removeValue != 5 {
		t.Errorf("Expected 5, got %d", *removeValue)
	}

	removeValue, _ = dll.RemoveAt(0)

	if *removeValue != 11 {
		t.Errorf("Expected 11, got %d", *removeValue)
	}

	length = dll.Len()

	if length != 0 {
		t.Errorf("Expected 0, got %d", length)
	}

	dll.Prepend(5)
	dll.Prepend(7)
	dll.Prepend(9)

	getValue, _ = dll.Get(2)

	if *getValue != 5 {
		t.Errorf("Expected 5, got %d", *getValue)
	}

	getValue, _ = dll.Get(0)

	if *getValue != 9 {
		t.Errorf("Expected 9, got %d", *getValue)
	}

	removeValue = dll.Remove(9)

	if *removeValue != 9 {
		t.Errorf("Expected 9, got %d", *removeValue)
	}

	length = dll.Len()

	if length != 2 {
		t.Errorf("Expected 2, got %d", length)
	}

	getValue, _ = dll.Get(0)

	if *getValue != 7 {
		t.Errorf("Expected 7, got %d", *getValue)
	}
}
