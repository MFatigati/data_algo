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

func sumCandidate(nums []int) int {
	sum := 0
	for _, elem := range nums {
		sum += elem
	}
	return sum
}

func backtrack(list []int, candidate []int, results *[][]int, target int) {
	if sumCandidate(candidate) > target {
		return
	}

	if sumCandidate(candidate) == target { // also can return if terminal condition (lead node)
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
		// fmt.Println("candidate: ", candidate)
		// fmt.Println("results: ", results)
		return
	}

	for i := 0; i < len(list); i++ {
		// if sumCandidate(candidate) + list[i] > target {
		// 	continue
		// }
		candidate = append(candidate, list[i]) // take
		fmt.Println("candidate: ", candidate)
		// you will go as many levels deep with this recursion as it takes to come to a base case
		// when you come to a base case, you will return from the lowest recursive call, and the next-lowest
		// recursive call will continue with its for loop.
		// When that next-lowest call finishes
		backtrack(list[i:], candidate, results, target) // explore
		// each time you return, you come back to pop off the last value and continue working through
		// additional possibilities via the current for loop
		// when the current for loop ends, how does the function know to return?

		candidate = candidate[:len(candidate)-1] // clean up
	}
}

func combinationSum(list []int, target int) [][]int {
	var results [][]int
	var candidate []int
	backtrack(list, candidate, &results, target)
	return results
}

func main() {

	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))

}
