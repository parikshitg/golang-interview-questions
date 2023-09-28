// Question 1: How do you stop a goroutine?
package main

import (
	"fmt"
	"time"
)

var count = 0

func routine(done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			count++
			fmt.Println(count)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	done := make(chan bool)
	go routine(done)
	time.Sleep(5 * time.Second)
	done <- true
}
