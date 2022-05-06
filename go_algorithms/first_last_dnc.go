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

func leftHalf(nums []int) []int {
	midpoint := len(nums) / 2
	return nums[0:midpoint]
}

func rightHalf(nums []int) []int {
	midpoint := len(nums) / 2
	return nums[midpoint:]
}

func resolveIndexes(leftResults []int, rightResults []int, target int, lengthLeft int) []int {
	significantIndexes := []int{-1, -1}
	// get most significant left index value
	if leftResults[0] != -1 {
		significantIndexes[0] = leftResults[0]
	} else if rightResults[0] != -1 {
		significantIndexes[0] = rightResults[0] + lengthLeft
	} else {
		significantIndexes[0] = -1
	}

	// get most significant right index value
	if rightResults[1] != -1 {
		significantIndexes[1] = rightResults[1] + lengthLeft
	} else if leftResults[1] != -1 {
		significantIndexes[1] = leftResults[1]
	} else {
		significantIndexes[1] = -1
	}

	return significantIndexes
}

func searchRange(nums []int, target int) []int {
	// base cases; very different for every problem
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	// divide
	left := leftHalf(nums)
	right := rightHalf(nums)
	// conquer (recurse)
	leftResult := searchRange(left, target)
	rightResult := searchRange(right, target)
	// combine; very different for every problem
	currentIndexes := resolveIndexes(leftResult, rightResult, target, len(left))

	return currentIndexes
}

func main() {

	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 8, 10}, 8)) // [3, 6]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))       // [3, 4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 10}, 8))    // [3, 5]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 10}, 7))    // [1, 2]
	fmt.Println(searchRange([]int{5, 7, 7, 7, 7, 7, 10}, 7))    // [1, 5]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))       // [-1, -1]
	fmt.Println(searchRange([]int{}, 0))                        // [-1, -1]
	fmt.Println(searchRange([]int{2, 2}, 2))                    // [0, 1]
	fmt.Println(searchRange([]int{2, 4}, 2))                    // [0, 0]
}
