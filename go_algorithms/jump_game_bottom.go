/*


At the end, return whether it is true for index 0

For the last index, value is true
Iterate back to index zero
For the index prior to that, if any of the values it can reach are true, it is true
	For the index prior to that, if any of the values it can reach are true, it is true
*/

package main

import "fmt"

func canJump(nums []int) bool {
	// create a map to state whether it is possible to reach the end from a given index
	canThisIndexMakeIt := map[int]bool{}
	lastIdx := len(nums) - 1
	// the value for the last index is always true
	canThisIndexMakeIt[lastIdx] = true

	// iterate over all the prior indexes until the beginning, starting at the second to last
	for i := lastIdx - 1; i >= 0; i-- {
		// check whether any of the steps reachable from the current index have a true value
		howManyStepsFromHere := nums[i]
		for step := 1; step <= howManyStepsFromHere; step++ {
			// if one of the reachable steps does have a true value, set the value for the current index to true
			if canThisIndexMakeIt[i+step] {
				canThisIndexMakeIt[i] = true
				break
			}
		}
	}
	fmt.Println(canThisIndexMakeIt)
	// return the value for whether the first index can make it
	return canThisIndexMakeIt[0]
}

func main() {

	fmt.Println(canJump([]int{2, 3, 1, 2, 4})) // true
	// fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
	// fmt.Println(canJump([]int{3, 2, 1, 1, 4})) // true
	// fmt.Println(canJump([]int{0}))             // true
	// fmt.Println(canJump([]int{2, 0, 0}))       // true
	// fmt.Println(canJump([]int{2, 5, 0, 0}))    // true

}
