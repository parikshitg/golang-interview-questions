// Question 6: How can you implement stacks in go?
// LIFO : Last IN First OUT
package main

import (
	"fmt"
)

type Stack []int

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() {
	if len(*s) <= 0 {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	s.Push(4)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}
