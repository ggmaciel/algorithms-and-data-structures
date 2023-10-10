package queue

import (
	"errors"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue[T]) Peek() (*T, error) {
	if q.head == nil {
		return nil, errors.New("Queue is empty")
	}
	return &q.head.value, nil
}

func (q *Queue[T]) Deque() (*T, error) {
	if q.head == nil {
		return nil, errors.New("Queue is empty")
	}
	q.length--
	head := q.head
	q.head = q.head.next

	if q.length == 0 {
		q.tail = nil
	}

	return &head.value, nil
}

func (q *Queue[T]) Enqueue(item T) {
	node := &Node[T]{value: item, next: nil}
	q.length++
	if q.length == 1 {
		q.tail = node
		q.head = node
		return
	}
	q.tail.next = node
	q.tail = node
}
