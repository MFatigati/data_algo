/*
An array is jumpable if any of the subarrays starting at the indexes reachable from its first value are jumpable arrays
*/

package main

import "fmt"

// func canJump(nums []int) bool {
// 	if helper(nums) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func helper(nums []int) bool {
// 	currentJumps := nums[0]
// 	subslices := [][]int{}
// 	for sub := 1; sub <= currentJumps && sub < len(nums); sub++ {
// 		subslices = append(subslices, nums[sub:])
// 	}
// 	fmt.Println(currentJumps, subslices)
// 	for sub := 0; sub < len(subslices); sub++ {
// 		canJumpFromFirst := subslices[sub][0]
// 		lengthToJump := len(subslices[sub]) - 1
// 		fmt.Println(canJumpFromFirst, lengthToJump)
// 		if canJumpFromFirst == lengthToJump {
// 			return true
// 		} else {
// 			return helper(subslices[sub])
// 		}
// 	}
// 	return false
// }

/*
if length is zero or one, return true
if current jumps possible from index 0 is greater than or equal to length to jump to the end, return true
if any of the subarrays reachable from the first index are jumpable, return true
otherwise, return false

*/

func canJump(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	currentJumps := nums[0]
	lengthToJump := len(nums) - 1
	if currentJumps >= lengthToJump {
		return true
	}

	subslices := [][]int{}
	for sub := 1; sub <= currentJumps && sub < len(nums); sub++ {
		subslices = append(subslices, nums[sub:])
	}

	for sub := 0; sub < len(subslices); sub++ {

		if canJump(subslices[sub]) {
			return true
		}
	}

	return false
}

func main() {

	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))                                                                                                                                                                                                                                                                                                          // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))                                                                                                                                                                                                                                                                                                          // false
	fmt.Println(canJump([]int{3, 2, 1, 1, 4}))                                                                                                                                                                                                                                                                                                          // true
	fmt.Println(canJump([]int{0}))                                                                                                                                                                                                                                                                                                                      // true
	fmt.Println(canJump([]int{2, 0, 0}))                                                                                                                                                                                                                                                                                                                // true
	fmt.Println(canJump([]int{2, 5, 0, 0}))                                                                                                                                                                                                                                                                                                             // true
	fmt.Println(canJump([]int{3, 0, 8, 2, 0, 0, 1}))                                                                                                                                                                                                                                                                                                    // true
	fmt.Println(canJump([]int{2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6})) // false

}
