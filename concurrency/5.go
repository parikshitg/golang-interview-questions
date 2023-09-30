// Question 5: What's wrong this this program.
package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))

	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

// any time a goroutine depends on a variable whose value might change,
// you must pass the value into the goroutine.

// for _, v := range a {
// 	go func(x int) {
// 		ch <- x * 2
// 	}(v)
// }
