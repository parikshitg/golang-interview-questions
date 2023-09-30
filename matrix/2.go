// Question 1: 2D matrix multiplication
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

	// m1 x m2  = {
	// {11, 11, 11},
	// {29, 29, 29},
	// {47, 47, 47},
	// }

	m3 := [3][3]int{}

	for i := range m1 {
		for j := range m2 {
			sum := 0
			for k := range m2 {
				sum += m1[i][k] * m2[k][j]
			}
			m3[i][j] = sum
		}
	}

	fmt.Println(m3)

	// -----------------------------
	m4 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	m5 := [2][3]int{
		{5, 6, 7},
		{8, 9, 10},
	}

	// m4 x m5 = {
	// 	{21, 24, 27},
	// 	{47, 54, 61},
	// }

	m6 := [2][3]int{}

	for i := range m4 {
		for j := range m5[0] {
			sum := 0
			for k := range m5 {
				sum += m4[i][k] * m5[k][j]
			}
			m6[i][j] = sum
		}
	}

	fmt.Println(m6)
}
