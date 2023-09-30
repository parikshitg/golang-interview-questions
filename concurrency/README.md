# Golang Concurrency Questions

1. How to stop a goroutine?
2. Implement the SumOfSquares function which takes an integer, n and returns the sum of all squares between 1 and n. You’ll need to use select statements, goroutines, and channels. For example, entering 5 would return 55 because (1)^2 + (2)^2 + (3)^2 + (4)^2 + (5)^2 = 55.
3. Using two goroutines one that prints odd numbers and one even numbers, print 1-10 in sequence. 
4. Read and write Fibonacci series to channel.
5. What's wrong this this program.
```
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
```
