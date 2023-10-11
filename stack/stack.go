package main

import "errors"

type Node[T any] struct {
	value T
	prev  *Node[T]
}

type Stack[T any] struct {
	length int
	head   *Node[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		length: 0,
		head:   nil,
	}
}

func (s *Stack[T]) Peek() (*T, error) {
	if s.head == nil {
		return nil, errors.New("Stack is empty")
	}
	return &s.head.value, nil
}

func (s *Stack[T]) Push(item T) {
	node := &Node[T]{value: item, prev: nil}

	s.length++
	if s.head == nil {
		s.head = node
		return
	}
	node.prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.head == nil {
		return nil, errors.New("Stack is empty")
	}

	s.length = max(0, s.length-1)
	if s.length == 0 {
		head := s.head
		s.head = nil
		return &head.value, nil
	}
	head := s.head
	s.head = s.head.prev
	return &head.value, nil
}
