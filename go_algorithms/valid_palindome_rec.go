/*
P
	I:
	O:
E
D
A
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(str string) bool {
	regStr := strings.Join(regexp.MustCompile("([[:alpha:]]|\\d)*").FindAllString(strings.ToLower(str), -1), "")
	fmt.Println(regStr)

	if len(regStr) == 1 || len(regStr) == 0 {
		return true
	}
	if regStr[0] != regStr[len(regStr)-1] {
		return false
	}
	return isPalindrome(regStr[1 : len(regStr)-1])
}

// alternate solution involves passing pointers
// if the first and last values are equal, and the rest of the string is a palidrome, return true

/*
isPalidrome("abcedcba")
	= isPalidrome("bcddcb")
		= isPalidrome("cddc")
			= isPalidrome("dd")
				= isPalidrome("d")
					= true

*/

func main() {

	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println(isPalindrome("race a car"))
	// fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("abcedcba"))
}
