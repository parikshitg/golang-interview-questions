// Question 8: Print pair of numbers in an array whose sum will result into a
// particular number, i.e. if array given [1,2,3,3,4,5] and sum given 6,
// then no. of pairs will be 3 -(1,5), (3,3), (2,4)
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 3, 4, 5}
	sum := 6

	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Println(arr[i], arr[j])
			}
		}
	}
}
