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

func backtrackSubs(collection []int, candidate []int, results *[][]int, position int) {
	copyCandidate := make([]int, len(candidate))
	copy(copyCandidate, candidate)
	*results = append(*results, copyCandidate)

	for i := position; i < len(collection); i++ {
		candidate = append(candidate, collection[i]) // take
		//fmt.Println(candidate)
		// call the backtrack function so that the next starting position is 1 greater than current starting position
		backtrackSubs(collection, candidate, results, position+1) // explore
		fmt.Println(candidate)
		candidate = candidate[:len(candidate)-1] // clean up
	}
}

func subsets1(nums []int) [][]int {
	var results [][]int
	var candidate []int
	backtrackSubs(nums, candidate, &results, 0)
	return results
}

func main() {
	//fmt.Println(uniqueSubset([]int{2, 3}, &[][]int{[]int{3, 2}}))
	fmt.Println(subsets1([]int{1, 2, 3}))
	// , []int{1, 2, 3}
}
