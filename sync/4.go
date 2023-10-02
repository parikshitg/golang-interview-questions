// Question 4: Race Condition
// When multiple threads try to access and modify same data, race condition occurs.
// For example a thread tries to update an integer while other thread tries to read it.
package main

// one goroutine is the default goroutine
import (
	"fmt"
	"runtime"
	"sync"
)

var wgIns sync.WaitGroup

func main() {
	counter := 0

	for i := 0; i < 50; i++ {
		wgIns.Add(1)
		go func() {
			for j := 0; j < 50; j++ {
				counter += 1
			}
			wgIns.Done()
		}()
	}

	fmt.Println("The number of goroutines before wait = ", runtime.NumGoroutine())
	wgIns.Wait()
	fmt.Println("Counter value = ", counter)
	fmt.Println("The number of goroutines after wait = ", runtime.NumGoroutine())
}

// The count of the number of goroutines before and after the wait is inconsistent.
// This  inconsistency is caused by racial factors.
