/*
P
	I:
	O:
E
D
A

Set currentBase index to 0
Set currentSub index to 0
While currentBase is less than base length - 1
	check if value at currentBase equals value at currentSub
	if so, advance currentSub index
	if currentSub index is greater than length of currentSub, return true
	increment currentBase
Return false

O(N)
*/

package main

import "fmt"

func isSubsequence(sub string, base string) bool {
	if sub == "" {
		return true
	}

	currentBaseI := 0
	currentSubI := 0

	for currentBaseI <= len(base)-1 {
		if base[currentBaseI] == sub[currentSubI] {
			fmt.Println(string(base[currentBaseI]), string(sub[currentSubI]))
			currentSubI++
			if currentSubI > len(sub)-1 {
				return true
			}
		}
		currentBaseI++
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
