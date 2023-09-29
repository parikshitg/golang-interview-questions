// Question 4: Read and write Fibonacci series to channel.
package main

import (
	"fmt"
)

func fibonacci(ch chan int, done chan bool) {
	x, y := 0, 1
	for {
		select {
		case <-done:
			close(ch)
			return
		default:
			ch <- x
			x, y = y, x+y
		}
	}
}

func main() {
	done := make(chan bool)
	defer close(done)
	ch := make(chan int)

	go fibonacci(ch, done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	done <- true
}
