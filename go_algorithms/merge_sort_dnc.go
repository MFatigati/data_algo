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

func mergeTwoSortedArrays(a []int, b []int) (result []int) {
	aPtr := 0
	bPtr := 0

	for aPtr < len(a) && bPtr < len(b) {
		if a[aPtr] < b[bPtr] {
			result = append(result, a[aPtr])
			aPtr++
		} else {
			result = append(result, b[bPtr])
			bPtr++
		}
	}

	if aPtr < len(a) {
		result = append(result, a[aPtr:]...)
	}
	if bPtr < len(b) {
		result = append(result, b[bPtr:]...)
	}

	return result
}

func mergeSort(list []int) []int {
	// base case
	if len(list) == 1 {
		return list
	}
	// divide
	left := leftHalf(list)
	right := rightHalf(list)
	// conquer (recurse)
	leftResult := mergeSort(left)
	rightResult := mergeSort(right)
	// combine; very different for every problem
	return mergeTwoSortedArrays(leftResult, rightResult)
}

func main() {
	fmt.Println(mergeSort([]int{3, 2, 5, 1, 6, 4}))
	//fmt.Println(mergeTwoSortedArrays([]int{1, 3}, []int{2}))
}
