/*

 */

package main

import "fmt"

func isPalindrome1(str string) bool {
	left := 0
	right := len(str) - 1

	for left <= right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func backtrack1(s string, candidate []string, results *[][]string) {
	if len(candidate) > 0 && !isPalindrome1(candidate[len(candidate)-1]) {
		return
	}
	if len(s) == 0 {
		copyCandidate := make([]string, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
		return
	}

	var temp string

	for i := 0; i < len(s); i++ {
		temp += string(s[i])
		candidate = append(candidate, string(temp))
		backtrack1(s[i+1:], candidate, results)
		candidate = candidate[:len(candidate)-1] // clean up
	}
}

func partition(s string) [][]string {
	var results [][]string
	var candidate []string
	backtrack1(s, candidate, &results)
	return results
}

func main() {

	fmt.Println(partition("aab"))
	//fmt.Println(isPalindrome1("aabbaa"))

}
