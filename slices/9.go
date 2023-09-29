// Question 9: Function to delete an slice element.
package main

import (
	"fmt"
)

func deleteArrayElement(arr *[]int, index int) {
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before deletion:", arr) // [1 2 3 4 5]
	deleteArrayElement(&arr, 2)          // delete element at index 2 i.e. 3
	fmt.Println("After deletion:", arr)  // [1 2 4 5]
}
