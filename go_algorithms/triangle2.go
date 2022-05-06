/*
starting at 0

base case:
row index = len triangle - 1
	return triangle[row index][column index]


shortest path currentIdx[0] = whichever is lesser of shortestPath(currentIdx[0]) || shortestPath(currentIdx[0 - 1])
if

currentIdx = len array - 1

*/

package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	memo := map[string]int{}
	return helper(triangle, 0, 0, memo)
}

func helper(triangle [][]int, row int, column int, memo map[string]int) int {
	fmt.Println(memo)
	if row == len(triangle)-1 {
		return triangle[row][column]
	} else {
		firstAdjacent := string(row+1) + ":" + string(column)
		secondAdjacent := string(row+1) + ":" + string(column+1)
		_, ok1 := memo[firstAdjacent]
		_, ok2 := memo[secondAdjacent]

		if !ok1 {
			memo[firstAdjacent] = helper(triangle, row+1, column, memo)
		}
		if !ok2 {
			memo[secondAdjacent] = helper(triangle, row+1, column+1, memo)
		}

		if (triangle[row][column] + memo[firstAdjacent]) < (triangle[row][column] + memo[secondAdjacent]) {
			return triangle[row][column] + memo[firstAdjacent]
		} else {
			return triangle[row][column] + memo[secondAdjacent]
		}
	}
}

func main() {

	triangle1 := [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}
	triangle2 := [][]int{[]int{-10}}
	triangle3 := [][]int{[]int{-1}, []int{2, 3}, []int{1, -1, -3}}

	fmt.Println(minimumTotal(triangle1))
	fmt.Println(minimumTotal(triangle2))
	fmt.Println(minimumTotal(triangle3))

}
