// Question 3. Print values of 2-D array in spiral format.
package main

import (
	"fmt"
)

func spiral(m [3][3]int) {
	top := 0
	bottom := len(m) - 1
	left := 0
	right := len(m) - 1

	for top <= bottom && left <= right {

		for i := left; i <= right; i++ {
			fmt.Print(m[top][i], " ")
		}
		top++

		for i := top; i <= bottom; i++ {
			fmt.Print(m[i][right], " ")
		}
		right--

		if top < bottom {
			for i := right; i >= left; i-- {
				fmt.Print(m[bottom][i], " ")
			}
		}
		bottom--

		if left < right {
			for i := bottom; i >= top; i-- {
				fmt.Print(m[i][left], " ")
			}
		}
		left++
	}
}

func main() {
	m := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	spiral(m) // 1, 2, 3, 6, 9, 8, 7, 4, 5
}
