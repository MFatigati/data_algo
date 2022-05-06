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

func intersection(nums1 []int, nums2 []int) []int {
	seen := map[int]int{}
	result := []int{}

	for _, elem := range nums1 {
		_, exists := seen[elem]
		if exists {
			continue
		} else {
			seen[elem] = 0
		}
	}

	for _, elem := range nums2 {
		_, exists := seen[elem]
		if exists {
			if seen[elem] == 0 {
				result = append(result, elem)
				seen[elem]++
			} else {
				continue
			}
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))    // [2]
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 4})) // [4, 9]
	fmt.Println(intersection([]int{3, 1, 1, 2}, []int{1, 1, 1})) // [1]
	fmt.Println(intersection([]int{}, []int{}))                  // []
	fmt.Println(intersection([]int{}, []int{5, 3, 1}))           // []
}
