/*
P
	I: array of numberes, nums
	O: nested arrays of numbers
		- each nested array contains a triplet
			- triplets must all be unique
			- triplets must add to zero
E
D
A

set first pointer to index zero
set second pointer to index one
set third pointer to index two

while first pointer is less than or equal to length minus 3
iterate over all with first
	iterate over all with second
		iterate over all with third

			if sum of all three is zero
					add triplet to return array

filter out duplicate triplets


*/

package main

import "fmt"

func threeSum(nums []int) (result [][]int) {

	for first := 0; first < len(nums)-3; first++ {
		for second := first + 1; second < len(nums)-2; second++ {
			for third := second + 1; third < len(nums)-1; third++ {
				currentSum := nums[first] + nums[second] + nums[third]
				// result = append(result, []int{first, second, third})
				if currentSum == 0 {
					result = append(result, []int{nums[first], nums[second], nums[third]})
				}
			}
		}
	}

	// FILTER OUT DUPLICATES SOMEHOW

	return result
}

func main() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{}))                    // []
	fmt.Println(threeSum([]int{}))                    // []
}
