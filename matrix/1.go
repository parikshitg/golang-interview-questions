// Question 1: 2D matrix addition
package main

import (
	"fmt"
)

func main() {

	m1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	m2 := [3][3]int{
		{2, 2, 2},
		{3, 3, 3},
		{1, 1, 1},
	}

	// m1 + m2  = {
	// {3, 4, 5},
	// {7, 8, 9},
	// {8, 9, 10},
	// }

	m3 := [3][3]int{}

	for i := range m1 {
		for j := range m2 {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}

	fmt.Println(m3)
}
