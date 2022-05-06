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

func uniqueSubset(candidate []int, results *[][]int) bool {
	if len(candidate) == 0 {
		return true
	}
	for _, result := range *results {
		matchCounter := 0
		for _, resultElem := range result {
			for _, candidateElem := range candidate {
				fmt.Println(candidateElem, resultElem, matchCounter)
				lengthCandidate := len(candidate)
				if resultElem == candidateElem {
					matchCounter++
					if matchCounter == lengthCandidate {
						return false
					}
				}
			}
		}
	}
	return true
}

func alreadyPresent(candidate []int, num int) bool {
	for _, elem := range candidate {
		if elem == num {
			return true
		}
	}
	return false
}

func backtrack2(collection []int, candidate []int, results *[][]int) {
	if !(uniqueSubset(candidate, results)) {
		return
	}
	if uniqueSubset(candidate, results) && len(*results) > 0 { // also can return if terminal condition (lead node)
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
		return
	}

	for i := 0; i < len(collection); i++ {
		if alreadyPresent(candidate, collection[i]) {
			continue
		}
		candidate = append(candidate, collection[i]) // take
		fmt.Println(candidate)
		backtrack2(collection[1:], candidate, results) // explore
		candidate = candidate[:len(candidate)-1]       // clean up
	}
}

func subsets(nums []int) [][]int {
	var results [][]int
	var candidate []int
	backtrack2(nums, candidate, &results)
	return results
}

func main() {
	//fmt.Println(uniqueSubset([]int{2, 3}, &[][]int{[]int{3, 2}}))
	fmt.Println(subsets([]int{1, 2, 3}))
	// , []int{1, 2, 3}
}
