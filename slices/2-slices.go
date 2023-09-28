// Question 2: Slices pass by value vs pass by reference.
package main

import (
	"fmt"
)

func addOneToSliceElements(slice []int) {
	for i := range slice {
		slice[i] = slice[i] + 1
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", slice)
	addOneToSliceElements(slice)
	fmt.Println("Original Slice after function call:", slice) // 2 3 4 5 6 as slices are passed by reference
}
