// Question 2: Arrays pass by value vs pass by reference.
package main

import (
	"fmt"
)

func addOneToArrayElements(array [5]int) {
	for i := range array {
		array[i] = array[i] + 1
	}
}

func squareOfArrayElements(array *[5]int) {
	for i := range array {
		array[i] = array[i] * array[i]
	}
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", array)
	addOneToArrayElements(array)
	fmt.Println("Original Array after function call:", array) // 1 2 3 4 5 as if you pass an array to a function, it will receive a copy of the array, not a pointer to it.

	squareOfArrayElements(&array)                                   // But this style isn't idiomatic go, use Slices instead
	fmt.Println("Original Array after passed by reference:", array) // 1 4 9 16 25
}
