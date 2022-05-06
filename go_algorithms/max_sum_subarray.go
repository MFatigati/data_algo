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
)

func getCross(left []int, right []int) int {
	leftMax := 0
	leftRunningTotal := 0
	leftPtr := len(left) - 1

	rightMax := 0
	rightRunningTotal := 0
	rightPtr := 0

	for leftPtr >= 0 {
		if leftPtr == len(left)-1 {
			leftRunningTotal = left[leftPtr]
			leftMax = left[leftPtr]
		} else {
			leftRunningTotal += left[leftPtr]
			if leftRunningTotal > leftMax {
				leftMax = leftRunningTotal
			}
		}
		leftPtr--
	}

	for rightPtr < len(right) {
		if rightPtr == 0 {
			rightRunningTotal = right[rightPtr]
			rightMax = right[rightPtr]
		} else {
			rightRunningTotal += right[rightPtr]
			if rightRunningTotal > rightMax {
				rightMax = rightRunningTotal
			}
		}
		rightPtr++
	}

	return rightMax + leftMax
}

func greatest(a int, b int, c int) int {
	greatest := 0
	if a < b {
		greatest = b
	} else {
		greatest = a
	}
	if c > greatest {
		greatest = c
	}
	return greatest
}

func leftHalf(nums []int) []int {
	midpoint := len(nums) / 2
	return nums[0:midpoint]
}

func rightHalf(nums []int) []int {
	midpoint := len(nums) / 2
	return nums[midpoint:]
}

func maxSubArray(nums []int) int {
	// base cases; very different for every problem
	if len(nums) == 1 {
		return nums[0]
	}
	// divide
	left := leftHalf(nums)
	right := rightHalf(nums)
	// conquer (recurse)
	leftResult := maxSubArray(left)
	rightResult := maxSubArray(right)
	// combine; very different for every problem
	currentCross := getCross(left, right)

	return greatest(leftResult, rightResult, currentCross)
}

func main() {

	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6: [4,-1,2,1]
	//fmt.Println(getCross([]int{0, 1, -2}, []int{3, -4, 5}))
}
