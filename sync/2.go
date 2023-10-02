// Questions 2: Waitgroups.
// To wait for multiple goroutines to finish we use waitgroups.
package main

import (
	"fmt"
	"sync"
)

func hello(s string, w *sync.WaitGroup) {
	fmt.Println("Hello", s)
	w.Done()
}

func main() {
	w := &sync.WaitGroup{}
	w.Add(1)
	go hello("World", w)
	w.Wait()
}
