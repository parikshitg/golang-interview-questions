// Question 1: How do you stop a goroutine?
// Using context to avoid leaking goroutines.
package main

import (
	"context"
	"fmt"
	"time"
)

var count = 0

func routine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			count++
			fmt.Println(count)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go routine(ctx)
	time.Sleep(5 * time.Second)
}
