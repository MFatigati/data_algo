/*
P

// Reverse Vowels of a String:
// https://leetcode.com/problems/reverse-vowels-of-a-string/
// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.
//
// Example 1:
// Input: s = "hello"
// Output: "holle"
//
// Example 2:
// Input: s = "leetcode"
// Output: "leotcede"
// "eOiau" => "uaiOe"
// "eoiauo" => "ouaioe"
// Constraints:
// 1 <= s.length <= 3 * 105
// s consist of printable ASCII characters.
P:
	I: a string of any ascii characters, in varying cases
	O: a string, with just the vowels reversed
	Rules:
		- keep capitalization with each particular letter
		-
E
"" => ""
D
	transform string into slice, then join it for the return
A
Split the string into a slice of all its values
Set a pointer to the first index
Set a second pointer to the last index

Create variables
	isPtr1onVowel bool
	isPtr2onVowel bool

helper function:
	isOnVowel

Until first pointer greater/equal to second pointer
	if ptr1 is not on a vowel, increment it
		if it is on a vowel, set isPtr1onVowel to true
	if ptr2 is not on a vowel, decrement it
		if it is on a vowel, set isPtr2onVowel to true
	if ptr1 and ptr2 are both on a vowel
			- check this by looking at the two boolean variables
		swap the vowels at those positions
		increment ptr1
		decrement ptr2
	set isPtr1onVowel/isPtr2onVowel to false

*/

package main

import (
	"fmt"
	"strings"
)

func isVowel(char string) bool {
	//re := regexp.MustCompile("aeiouAEIOU")
	vowels := map[string]bool{}
	vowels["a"] = true
	vowels["e"] = true
	vowels["i"] = true
	vowels["o"] = true
	vowels["u"] = true
	vowels["A"] = true
	vowels["E"] = true
	vowels["I"] = true
	vowels["O"] = true
	vowels["U"] = true
	_, ok := vowels[char]
	return ok
}

func reverseVowels(s string) (result string) {
	split := []string{}
	for _, elem := range s {
		split = append(split, string(elem))
	}

	ptr1 := 0
	ptr2 := len(s) - 1

	var isPtr1onVowel bool = false
	var isPtr2onVowel bool = false

	for ptr1 <= ptr2 {
		if !isVowel(split[ptr1]) {
			ptr1++
		} else {
			isPtr1onVowel = true
		}
		if !isVowel(split[ptr2]) {
			ptr2--
		} else {
			isPtr2onVowel = true
		}
		if isPtr1onVowel && isPtr2onVowel {
			split[ptr1], split[ptr2] = split[ptr2], split[ptr1]
			ptr1++
			ptr2--
			isPtr1onVowel = false
			isPtr2onVowel = false
		}
	}

	result = strings.Join(split, "")
	return result
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels(""))
}
