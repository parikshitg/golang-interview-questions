// Question 1: Mutex
package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

func incrementCounter(writer string, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	counter++
	fmt.Println(writer, ":", counter)
}

func main() {
	m := &sync.Mutex{}
	for i := 0; i < 3; i++ {
		go incrementCounter("ONE", m)
	}
	for i := 0; i < 3; i++ {
		go incrementCounter("TWO", m)
	}
	for i := 0; i < 3; i++ {
		go incrementCounter("THREE", m)
	}
	time.Sleep(time.Second)
}
