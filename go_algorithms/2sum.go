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

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for end > start {

		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		}

		if numbers[start]+numbers[end] > target {
			end--
		} else {
			start++
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [1, 2]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // [1, 3]
}
