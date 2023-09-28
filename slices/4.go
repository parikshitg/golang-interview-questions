// Question 4: Predict the output

package main

import (
	"fmt"
)

func main() {
	var a [5]int

	fmt.Println("a:", a)       // [0 0 0 0 0]
	fmt.Println("a[0]:", a[0]) // [0]
	a[4] = 4
	fmt.Println(a)                 // [0 0 0 0 4]
	fmt.Println("len(a):", len(a)) // 5
}
