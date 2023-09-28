// Question 2. Implement the SumOfSquares function which takes an integer,
// n and returns the sum of all squares between 1 and n. You’ll need to use
// select statements, goroutines, and channels. For example, entering 5 would return 55
// because (1)^2 + (2)^2 + (3)^2 + (4)^2 + (5)^2 = 55.

package main

import (
	"fmt"
)

func square(ch chan int, n int) {
	ch <- n * n
}

func sumOfSquares(n int) int {
	sum := 0
	ch := make(chan int)
	for i := 1; i <= n; i++ {
		go square(ch, i)
		sum += <-ch
	}
	return sum
}

func main() {
	n := 5
	fmt.Println(sumOfSquares(n))
}
