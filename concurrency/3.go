// Question 3: Using two goroutines one that prints odd numbers and one even numbers, print 1-10 in sequence.
package main

import (
	"fmt"
)

func odd(ch chan int, done chan bool) {
	for v := range ch {
		if v%2 != 0 {
			fmt.Println(v)
			done <- true
		} else {
			ch <- v
		}
	}
}

func even(ch chan int, done chan bool) {
	for v := range ch {
		if v%2 == 0 {
			fmt.Println(v)
			done <- true
		} else {
			ch <- v
		}
	}
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go even(ch, done)
	go odd(ch, done)

	for i := 1; i <= 10; i++ {
		ch <- i
		<-done
	}
}
