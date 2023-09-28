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