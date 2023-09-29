// Question 10. Remove duplicate values in an array.
package main

import (
	"fmt"
)

func removeDuplicates(arr []int) []int {
	var result []int

	m := make(map[int]struct{})
	for _, v := range arr {
		_, ok := m[v]
		if !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 1, 1, 2, 2, 2, 2, 3, 6, 7, 5, 6, 8, 7, 9, 9}
	fmt.Println(removeDuplicates(arr))
}
