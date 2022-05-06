/*
P
Write a binary search program to find a number in an array of sorted numbers from scratch

	I:
	O:
E
D
A

if the midpoint is equal to the target, return true
else, if midpoint

*/

package main

import "fmt"

func binarySearch(nums []int, target int) bool {
	lower := 0
	upper := len(nums) - 1
	for lower <= upper {
		midpoint := upper - lower/2
		if nums[midpoint] == target {
			return true
		}
		if nums[midpoint] < target {
			lower = midpoint + 1
		}
		if nums[midpoint] > target {
			upper = midpoint - 1
		}
	}
	return false
}

func main() {

	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 2))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
}
