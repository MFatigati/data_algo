/*
find the first occurence of some substring in a base string
return the starting index of the substring within the base string


iterate over base string
	iterate over substring
		if not a match, break out of inner loop
	if reach a match
		- end outer loop
		- return result

iterate over base string until index plus length of substring is greater than length of base string
	get chunk that is the length of the substring
	if chunk matches substring, return current index

O(N), avoids double iteration

*/

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	i := 0

	for i+len(needle) < len(haystack)+1 {
		var chunk string = haystack[i : i+len(needle)]
		if chunk == needle {
			return i
		}
		i++
	}

	return -1
}

func main() {

	fmt.Println(strStr("lsdjflaabccjsdakjflaskdjaabbccfdsflsdaf", "aabbcc")) // 24
	fmt.Println(strStr("hello", "ll"))                                       // 2
	fmt.Println(strStr("", ""))                                              // 0
	fmt.Println(strStr("aaaaa", "bba"))                                      // -1

}
