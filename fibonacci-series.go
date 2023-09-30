package main

import (
	"fmt"
)

func recursive(n int) int {
	if n < 2 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}

func sequential(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	n := 10

	for i := 0; i < n; i++ {
		fmt.Print(recursive(i), " ")
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Print(sequential(i), " ")
	}
	fmt.Println()
}
