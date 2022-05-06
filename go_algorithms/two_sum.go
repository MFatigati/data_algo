/*
P
	I: array of integers: nums, integer: target
	O: slice of two integers
E
D
A

Create an empty map
Iterate over the way

*/

package main

import "fmt"

func twoSum(nums []int, target int) (indexes []int) {
	seen := map[int]int{}

	for i, elem := range nums {
		necessary := target - elem
		priorIdx, exists := seen[necessary]
		if exists {
			indexes = []int{priorIdx, i}
		} else {
			seen[elem] = i
		}
	}
	return indexes
}

func main() {

	fmt.Println(twoSum([]int{3, 3}, 6))
}
