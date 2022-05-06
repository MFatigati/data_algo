/*
P
	I:
	O:
E
D
A

Set start at index zero
set end at index length minus 1
	while end is greater or equal to start
		swap values
		increment start
		decrement end

*/

package main

import (
	"fmt"
)

func reverseString(s []byte) {
	start := 0
	end := len(s) - 1

	for end >= start {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func main() {
	fmt.Println(([]byte("hello")))
}
