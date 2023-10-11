package main

import "fmt"

func main() {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	popped1, err := stack.Pop()
	fmt.Printf("Popped: %v, Error: %v\n", *popped1, err) // Expected 3 and nil Error

	popped2, err := stack.Pop()
	fmt.Printf("Popped: %v, Error: %v\n", *popped2, err) // Expected 2 and nil Error

	top, err := stack.Peek()
	fmt.Printf("Top: %v, Error: %v\n", *top, err) // Expected 1 and nil Error

	popped3, err := stack.Pop()
	fmt.Printf("Popped: %v, Error: %v\n", *popped3, err) // Expected 1 and nil Error

	popped4, err := stack.Pop()
	fmt.Printf("Popped: %v, Error: %v\n", popped4, err) // Expected Error

	top2, err := stack.Peek()
	fmt.Printf("Top: %v, Error: %v\n", top2, err) // Expected Error
}
