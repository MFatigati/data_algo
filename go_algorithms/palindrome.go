package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	filtered := ""
	for _, elem := range s {
		if unicode.IsLetter(elem) || unicode.IsNumber(elem) {
			filtered += strings.ToLower(string(elem))
		}
	}
	var leftIdx int = 0
	var rightIdx int = len(filtered) - 1
	fmt.Println(leftIdx, rightIdx)
	for rightIdx >= leftIdx {
		if filtered[leftIdx] != filtered[rightIdx] {
			return false
		}
		rightIdx--
		leftIdx++
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("0P"))
}
