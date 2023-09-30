// Question 1. If given second as input then print the time duration of the
// input in hour, minute and second, e.g. if input is 318 then output will be 5m18s.
package main

import (
	"fmt"
)

func time2String(n int) string {
	seconds := n % 60
	minutes := n / 60
	if minutes < 60 {
		return fmt.Sprintf("%dm%ds", minutes, seconds)
	}
	hours := minutes / 60
	return fmt.Sprintf("%dh%dm%ds", hours, minutes, seconds)
}

func main() {
	fmt.Println(time2String(31800))
}
