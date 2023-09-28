// Question 1: How do you reverse a slice?
package main

import (
	"fmt"
)

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Slice:", slice)
	reverseSlice(slice)
	fmt.Println("Reversed Slice:", slice)
}
