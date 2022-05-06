/*

Create a map to hold the integer values for the sums at various indexes

For all the values in the last subarray, set their index values to their element values (no computation)

Iterate over all the remaining subarrays, unti lthe first
	For each index, store itself, plus the lesser of its two adjacent indexes

return DP[0,0]
*/

package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	totals := map[[2]int]int{}
	lastRow := len(triangle) - 1

	for i := 0; i < len(triangle[lastRow]); i++ {
		totals[[2]int{lastRow, i}] = triangle[lastRow][i]
	}

	for row := len(triangle) - 1; row >= 0; row-- {
		currentRow := triangle[row]
		for j := 0; j < len(currentRow); j++ {
			if totals[[2]int{row + 1, j}] < totals[[2]int{row + 1, j + 1}] {
				totals[[2]int{row, j}] = currentRow[j] + totals[[2]int{row + 1, j}]
			} else {
				totals[[2]int{row, j}] = currentRow[j] + totals[[2]int{row + 1, j + 1}]
			}
		}
	}

	return totals[[2]int{0, 0}]
}

func main() {
	triangle1 := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle1))

}
