/*
P
	I: array of integers, sorted ascending
	O: array: [first occurence of target value, last occurence of target value]
E
D
A

To get upper
Set lower to 0
Set upper to length nums minus 1
Set midpoint to upper - lower / 2
if midpoint is less than target, set upper to mid - 1
if midpoint is greater than target, set lower to mid + 1
if midpoint = target
	check to see if midpoint is FOLLOWED by another occurence of target
		if it IS followed by another occurence of the target, set lower to midpoint + 1
	if NOT followed by another occurence of target, set upper idx to current midpoint index

To get lower
Set lower to 0
Set upper to length nums minus 1
Set midpoint to upper - lower / 2
if midpoint is less than target, set upper to mid - 1
if midpoint is greater than target, set lower to mid + 1
if midpoint = target
	check to see if midpoint is PRECEDED by another occurence of the target
		if it IS preceded by another occurence of the target, set upper to midpoint - 1
	if NOT preceded by another occurence of target, set lower idx to current midpoint index
*/

package main

import "fmt"

func getLowerBinary(nums []int, target int) int {
	lower := 0
	upper := len(nums) - 1
	for lower <= upper {
		midpoint := (upper + lower) / 2
		//fmt.Println(lower, upper, midpoint, nums[midpoint])
		if nums[midpoint] < target {
			lower = midpoint + 1
		}
		if nums[midpoint] > target {
			upper = midpoint - 1
		}
		if nums[midpoint] == target {
			if midpoint == 0 {
				return midpoint
			}
			if nums[midpoint-1] == target {
				upper = midpoint - 1
			} else {
				return midpoint
			}
		}
	}
	return -1
}

func getUpperBinary(nums []int, target int) int {
	lower := 0
	upper := len(nums) - 1
	for lower <= upper {
		midpoint := (upper + lower) / 2
		//fmt.Println(lower, upper, midpoint, nums[midpoint])
		if nums[midpoint] < target {
			lower = midpoint + 1
		}
		if nums[midpoint] > target {
			upper = midpoint - 1
		}
		if nums[midpoint] == target {
			if midpoint == len(nums)-1 {
				return midpoint
			}
			if nums[midpoint+1] == target {
				lower = midpoint + 1
			} else {
				return midpoint
			}
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	lower := getLowerBinary(nums, target)
	upper := getUpperBinary(nums, target)
	return []int{lower, upper}
}

func main() {

	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 8, 10}, 8)) // [3, 4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))       // [-1, -1]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 10}, 8))    // [3, 5]
	fmt.Println(searchRange([]int{}, 0))                        // [-1, -1]
	fmt.Println(searchRange([]int{2, 2}, 2))                    // [0, 1]
	fmt.Println(searchRange([]int{2, 4}, 2))                    // [0, 0]

}
