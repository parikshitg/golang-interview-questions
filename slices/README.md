# Golang Slices Questions

1. How can you reverse a slice/array?
2. Pass by value and pass by reference of slice/arrays.
3. Arrays vs Slices. 
4. Predict the output
```
package main

import (
	"fmt"
)

func main() {
	var a [5]int

	fmt.Println("a:", a)
	fmt.Println("a[0]:", a[0])
	a[4] = 4
	fmt.Println(a)
	fmt.Println("len(a):", len(a))
}

```

5. Predict the output
```
package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 0, 4)
	fmt.Println(arr)

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 6)
	fmt.Println(arr)

	brr := arr[0:3]
	brr = append(brr, 7)
	fmt.Println(brr)

	a := make([]int, 0, 20)
	bah(a)
	fmt.Println(a)
}

func bah(a []int) {
	a = append(a, 10)
}
```

6. How to implement Stack in golang?
7. How to implement Queue in golang?
8. Print pair of numbers in an array whose sum will result into a particular number, i.e. if array given [1,2,3,3,4,5] and sum given 6, then no. of pairs will be 3 -(1,5), (3,3), (2,4)
9. Function to delete an slice element.
10. Remove duplicate values in an array.
