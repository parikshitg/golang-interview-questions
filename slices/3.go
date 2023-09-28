// ARRAYS
// In Golang: Arrays are values.

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	var b []int
	b = a // Assigning one array to other copies all the elements.
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// SLICES
	// Slices hold references to an underlying array
	// if you assign one slice to another, both refer to the same array.

}
