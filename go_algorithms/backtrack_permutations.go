/*

 */

// candidate     -> current path of exploration
// results  -> successful finds
package main

import "fmt"

func permute(nums []int) [][]int {
	var results [][]int
	var candidate []int
	backtrack(nums, candidate, &results)
	return results
}

func backtrack(list []int, candidate []int, results *[][]int) {
	if len(list) == 0 {
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
		return
	}

	for i := 0; i < len(list); i++ {
		candidate = append(candidate, list[i]) // take
		newList := filter(list, list[i])
		backtrack(newList, candidate, results)   // explore
		candidate = candidate[:len(candidate)-1] // clean up
	}
}

func filter(toFilter []int, num int) []int {
	newNums := []int{}
	for _, elem := range toFilter {
		if elem != num {
			newNums = append(newNums, elem)
		}
	}
	return newNums
}

func main() {

	fmt.Println(permute([]int{1, 2, 3}))
}
