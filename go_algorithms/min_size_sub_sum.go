/*
P
	I:
	O:
E
D
A



Set anchor at zero
set runner at zero
Set sum to zero
set minlength to 0

iterate over each index with anchor
	set runner to anchor
	set sum to zero
		while runner is less/equal to length nums - 1
			add value at runner to sum
			if sum is greater/equal to target
				if length of subarray is less than minlength
					set minlength to currentlength
					break
			advance runner

pretty sure this is ON^2, and not sure how to fix that

set anchor at zero
set runner at zero
set sum to zero
set minlength to zero
while runner is less than len nums minus 1
	get sum for current slice
	if sum is less than target
		advance runner
	if sum is greater than or equal to target
		set minlength to length of current sub
		reset sum to zero
		advance anchor to runner
end if result is 1, or second runner is out of bounds
*/

package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	minlength := 0
	anchor := 0
	runner := 0
	sum := 0

	for runner < len(nums) && anchor <= runner {
		fmt.Println(nums[anchor:runner+1], minlength, sum)
		if sum < target {
			sum += nums[runner]
			runner++
		} else if sum >= target {
			if minlength == 0 {
				minlength = runner - anchor
			} else {
				if runner-anchor+1 < minlength {
					minlength = runner - anchor
				}
			}
			sum -= nums[anchor]
			// if anchor != runner {
			// 	anchor++
			// }
			anchor++
		}
	}
	return minlength
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})) // 2
	// fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	// fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
	// fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))          // 3

}

// O(N^2) solution
// func minSubArrayLen(target int, nums []int) int {
// 	minlength := 0

// 	for anchor := 0; anchor < len(nums); anchor++ {
// 		sum := 0
// 		for runner := anchor; runner < len(nums); runner++ {
// 			sum += nums[runner]
// 			if sum >= target {
// 				if minlength == 0 {
// 					minlength = runner - anchor + 1
// 				} else {
// 					if runner-anchor+1 < minlength {
// 						minlength = runner - anchor + 1
// 					}
// 				}
// 				break
// 			}
// 		}
// 	}
// 	return minlength
// }
