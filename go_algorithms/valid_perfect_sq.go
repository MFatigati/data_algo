/*
P
	I:
	O:
E
D
A
*/

package main

import "fmt"

func isPerfectSquare(num int) bool {
	lower := 0
	upper := num
	midpoint := upper - lower/2

	for lower <= upper {
		midpoint = upper + lower/2
		if midpoint*midpoint == num {
			return true
		} else {
			if midpoint*midpoint > num {
				upper = midpoint - 1
			} else {
				lower = midpoint + 1
			}
		}
	}
	return false

}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(25))
	fmt.Println(isPerfectSquare(26))
	fmt.Println(isPerfectSquare(2147483647))

}
