/*

set triangleRowsIdx to zero
set singleRowIdx to zero
set sum to zero
while triangleRowsIdx is less than input array length
	get index at index, or index plus 1 (if that exists)
	whichever is greater, add that to the sum, and set singleRowIdx to that number

-------------------

minimum path value for one row equals whatever value can connect to,
out of all the possible paths from the previous triangles,
the path that, together with it, results in the minimum

So each row needs access to all the possible minimum paths from the row before it,
to be able to choose the one that can add to one of its values and get the smallest value

if row = 0, return row[0]

For index zero of the largest row, look at all the possible paths
	Store the path as the key in the memo table, and the sum as the value

store the sum of all the paths from all the indexes
{
	0, 0, 0: 0[0],
	1, 0, 0: 1[0], 
	1, 1, 0: 1[1],
	2, 0, 0: 2[0] + 1[0]
	2, 0, 1: 
	2, 1, 1:
	2, 1, 2
	2, 2

	0[0]: 0[0]
	1[0]: 1[0] + 0[0]
	1[1]: 1[1] + 0[0]
	2[0]: 2[0] + 1[0]
	2[0]: 2[0] + 1[1]
	2[1]: 2[1] + 1[]


}

*/

package main

import (
	"fmt"
	"strconv"
)

func minimumTotal(triangle [][]int) int {
	memo := map[[][]int]int
	return helper(triangle, memo)
}

func helper(triangle [][]int, memo map[string]int) (sum int) {
	if len(triangle) == 0 {
		memo[strconv.Itoa(triangle[0][0])] = triangle[0][0]
	}

	memo[strconv.Itoa(triangle[0][0])] = triangle[0][0]
	return sum

	/*
	For index zero of current row, its least possible path
		is itself, plus whichever of its two prior connecting elements has the least possible path
		
	minimumTotal(index0) = index0 + whichever is less minimumtotal(0 in prior row) or minimumtotal(1 in prior row)
	*/
}

func main() {
	triangle1 := [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}
	triangle2 := [][]int{[]int{-10}}
	triangle3 := [][]int{[]int{-1}, []int{2, 3}, []int{1, -1, -3}}

	fmt.Println(minimumTotal(triangle1))
	// fmt.Println(minimumTotal(triangle2))
	// fmt.Println(minimumTotal(triangle3))
}

// triangleRowsIdx := 0
// singleRowIdx := 0
// sum := 0
// for triangleRowsIdx < len(triangle) {

// 	if singleRowIdx+1 < len(triangle[triangleRowsIdx]) {
// 		if triangle[triangleRowsIdx][singleRowIdx] <
// 			triangle[triangleRowsIdx][singleRowIdx+1] {
// 			sum += triangle[triangleRowsIdx][singleRowIdx]
// 		}
// 		if triangle[triangleRowsIdx][singleRowIdx+1] <
// 			triangle[triangleRowsIdx][singleRowIdx] {
// 			sum += triangle[triangleRowsIdx][singleRowIdx+1]
// 			singleRowIdx++
// 		}
// 	} else {
// 		sum += triangle[triangleRowsIdx][singleRowIdx]
// 	}
// 	fmt.Println("row in triangle", triangleRowsIdx, "idx within row", singleRowIdx, "sum", sum)
// 	triangleRowsIdx++
// }
// return sum


/*
starting at 0

base case:
row index = len triangle - 1
	return triangle[row index][column index]


shortest path currentIdx[0] = whichever is lesser of shortestPath(currentIdx[0]) || shortestPath(currentIdx[0 - 1])
if 

currentIdx = len array - 1

*/