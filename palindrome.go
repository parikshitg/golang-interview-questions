package main

import (
	"fmt"
)

func isStringPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func isIntegerPalindrome(n int) bool {
	sum := 0
	input := n
	for n > 0 {
		r := n % 10
		sum = (sum * 10) + r
		n = n / 10
	}

	if input == sum {
		return true
	}
	return false
}

func main() {
	fmt.Println(isStringPalindrome("a"))       // true
	fmt.Println(isStringPalindrome("abcd"))    // false
	fmt.Println(isStringPalindrome("abcdcba")) // true
	fmt.Println(isStringPalindrome("racecar")) // true
	fmt.Println(isIntegerPalindrome(1))        // true
	fmt.Println(isIntegerPalindrome(12))       // false
	fmt.Println(isIntegerPalindrome(12321))    // true
}
