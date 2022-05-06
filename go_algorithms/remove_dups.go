/*
P
	I:
	O:
E
D
A

set anchor to 0
set runner to 1
while runner is less than length nums
	if nums[runner] does not equal nums[anchor]
		advance anchor by one
		replace value at anchor with value at runner (leave runner value as is)
	increment runner
return anchor index
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	anchor := 0
	runner := 1

	for runner < len(nums) {
		if nums[runner] != nums[anchor] {
			anchor++
			nums[anchor] = nums[runner]
		}
		runner++
	}
	return anchor + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))                      // 2, [1,2,_]
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})) // 5, [0,1,2,3,4,_,_,_,_,_]
}
