// Question 7: How to implement queue in go?
// FIFO: First IN first OUT
package main

import (
	"fmt"
)

type Queue []int

func NewQueue() Queue {
	return Queue([]int{})
}

func (q *Queue) Enqueue(n int) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() {
	if len(*q) <= 0 {
		return
	}
	*q = (*q)[1:]
}

func main() {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q)

	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
