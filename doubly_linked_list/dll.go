package doublylinkedlist

import "errors"

type Node[T comparable] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func NewDoublyLinkedList[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (dll *DoublyLinkedList[T]) Prepend(item T) {
	node := &Node[T]{value: item, next: nil, prev: nil}

	dll.length++
	if dll.head == nil {
		dll.head, dll.tail = node, node
		return
	}

	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}

func (dll *DoublyLinkedList[T]) InsertAt(item T, idx int) error {
	if idx > dll.length {
		return errors.New("index is out of bounds")
	} else if idx == dll.length {
		dll.Append(item)
		return nil
	} else if idx == 0 {
		dll.Prepend(item)
		return nil
	}

	dll.length++
	curr, err := dll.getAt(idx)

	if err != nil {
		return err
	}

	node := &Node[T]{value: item, next: curr, prev: curr.prev}
	curr.prev = node
	node.prev.next = node

	return nil
}

func (dll *DoublyLinkedList[T]) Append(item T) {
	dll.length++
	node := &Node[T]{value: item, next: nil, prev: nil}

	if dll.tail == nil {
		dll.head, dll.tail = node, node
		return
	}

	node.prev = dll.tail
	dll.tail.next = node
	dll.tail = node
}

func (dll *DoublyLinkedList[T]) Remove(item T) *T {
	curr := dll.head
	for i := 0; i < dll.length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return nil
	}

	return dll.removeNode(curr)
}

func (dll *DoublyLinkedList[T]) Get(idx int) (*T, error) {
	node, err := dll.getAt(idx)

	if err != nil {
		return nil, err
	}

	return &node.value, nil
}

func (dll *DoublyLinkedList[T]) RemoveAt(idx int) (*T, error) {
	if idx > dll.length {
		return nil, errors.New("index is out of bounds")
	}

	node, err := dll.getAt(idx)

	if err != nil {
		return nil, err
	}

	return dll.removeNode(node), nil
}

func (dll *DoublyLinkedList[T]) Len() int {
	return dll.length
}

func (dll *DoublyLinkedList[T]) removeNode(node *Node[T]) *T {
	dll.length--
	if dll.length == 0 {
		output := dll.head.value
		dll.tail, dll.head = nil, nil
		return &output
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == dll.head {
		dll.head = node.next
	}

	if node == dll.tail {
		dll.tail = node.prev
	}

	node.prev, node.next = nil, nil

	return &node.value
}

func (dll *DoublyLinkedList[T]) getAt(idx int) (*Node[T], error) {
	curr := dll.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	return curr, nil
}
