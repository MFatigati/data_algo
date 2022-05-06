package main

import (
	"fmt"
	"strconv"
)

func minimumTotal(triangle [][]int) int {
	memo := map[string]int{}
	return helper(triangle, 0, 0, memo) // this will get us the minimum sum path to the top
}

// helper takes a row and a column number, which is how we move down the triangle
func helper(triangle [][]int, row int, column int, memo map[string]int) int {
	// get individual value at current index
	currentVal := triangle[row][column]
	// base case, last row/subarray
	if row == len(triangle)-1 {
		return currentVal
	} else {
		// save adjacent index values in variables
		firstAdjacent := strconv.Itoa(row+1) + ":" + strconv.Itoa(column)
		secondAdjacent := strconv.Itoa(row+1) + ":" + strconv.Itoa(column+1)
		// check to see if values for adjacent indexes have already been computed
		_, ok1 := memo[firstAdjacent]
		_, ok2 := memo[secondAdjacent]

		// if not, compute them and store them in the memo
		if !ok1 {
			memo[firstAdjacent] = helper(triangle, row+1, column, memo)
			fmt.Println(memo)
		}
		if !ok2 {
			memo[secondAdjacent] = helper(triangle, row+1, column+1, memo)
			fmt.Println(memo)
		}

		// return value at current index, plus
		// the lesser of the values for the adjacent indexes in the memo table
		if (currentVal + memo[firstAdjacent]) < (currentVal + memo[secondAdjacent]) {
			return currentVal + memo[firstAdjacent]
		} else {
			return currentVal + memo[secondAdjacent]
		}
	}
}

func main() {

	triangle1 := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	// triangle2 := [][]int{{-10}}
	// triangle3 := [][]int{{-1}, {2, 3}, {1, -1, -3}}

	fmt.Println(minimumTotal(triangle1))
	// fmt.Println(minimumTotal(triangle2))
	// fmt.Println(minimumTotal(triangle3))

}
