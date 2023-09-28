// Question 5: Predict the output.
package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 0, 4)
	fmt.Println(arr) // []

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 6)
	fmt.Println(arr) // [1 2 3 6]

	brr := arr[0:3]
	brr = append(brr, 7)
	fmt.Println(arr) // [1 2 3 7]
	fmt.Println(brr) // [1 2 3 7]

	brr[1] = 100
	fmt.Println(arr) // 1 100 3 7
	fmt.Println(brr) // 1 100 3 7

	a := make([]int, 0, 20)
	bah(a)
	fmt.Println(a) // []
}

func bah(a []int) {
	a = append(a, 10)
}
