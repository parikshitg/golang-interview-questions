// Question 1: How do you reverse a array?
package main

import (
	"fmt"
)

func reverseArray(array *[10]int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Array:", array)
	reverseArray(&array)
	fmt.Println("Reversed Array:", array)
}
