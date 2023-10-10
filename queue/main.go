package main

import (
	"fmt"
	queue "queue/queue_package"
)

func main() {
	newQueue := queue.NewQueue[int]()

	_, err := newQueue.Peek()
	fmt.Println(err) // Expected 'Queue is empty'

	newQueue.Enqueue(1)

	next, _ := newQueue.Peek()
	fmt.Println(*next) // Expected 1

	removedValue, _ := newQueue.Deque()
	fmt.Println(*removedValue) // Expected 1

	_, err = newQueue.Peek()
	fmt.Println(err) // Expected 'Queue is empty'

	newQueue.Enqueue(2)
	newQueue.Enqueue(3)

	next, _ = newQueue.Peek()
	fmt.Println(*next) // Expected 2

	newQueue.Deque()

	next, _ = newQueue.Peek()
	fmt.Println(*next) // Expected 3
}
