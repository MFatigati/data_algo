package main

import "fmt"

/*
while there are nested arrays
iterate over outer array nested array
	set direction to forward
	iterate over each nested array
	if position is zero or last
		if its position is zero, add each element to new array in order, and remove from original array
		if its position is last, add each element to new array in reverse order, and remove from original array
	else
		if direction is forward, add last element onto new array, and remove from original array
		if direction is backward, add first element onto new array, and remove from original array
	remove any arrays of zero length
	reverse direction

N = number of nested arrays
M = elements in nested arrays
O(N*M)

recursively pass trimmed matrix
add top row, trim, and rotate
pointers: stop when bottom is greater than top, or left is greater than right
	- or, stop when you have the same number of elements in the output as in the matrix

*/

func appendAllInOrderThenRemove(current *[]int, result *[]int) {
	copy := *current
	for i := 0; i < len(*current); i++ {
		newValue := copy[i]
		*result = append(*result, newValue)
	}
	*current = []int{}
}

func appendAllInReverseOrderThenRemove(current *[]int, result *[]int) {
	copy := *current
	for i := len(*current) - 1; i >= 0; i-- {
		newValue := copy[i]
		*result = append(*result, newValue)
	}
	*current = []int{}
}

func appendLastThenRemove(current *[]int, result *[]int) {
	copy := *current
	*result = append(*result, copy[len(copy)-1])
	*current = copy[:len(copy)-1]
}

func appendFirstThenRemove(current *[]int, result *[]int) {
	copy := *current
	*result = append(*result, copy[0])
	*current = copy[1:]
}

func spiralOrder(matrix [][]int) (result []int) {
	direction := "forward"
	for len(matrix) > 0 {

		if direction == "forward" {
			for i := 0; i < len(matrix); i++ {
				if i == 0 || i == len(matrix)-1 {
					if i == 0 {
						appendAllInOrderThenRemove(&matrix[i], &result)
					}
					if i == len(matrix)-1 {
						appendAllInReverseOrderThenRemove(&matrix[i], &result)
					}
				} else {
					appendLastThenRemove(&matrix[i], &result)
				}
			}
		} else if direction == "backward" {
			for i := len(matrix) - 1; i >= 0; i-- {
				if i == 0 || i == len(matrix)-1 {
					if i == 0 {
						appendAllInOrderThenRemove(&matrix[i], &result)
					}
					if i == len(matrix)-1 {
						appendAllInReverseOrderThenRemove(&matrix[i], &result)
					}
				} else {
					appendFirstThenRemove(&matrix[i], &result)
				}
			}
		}

		temp := [][]int{}

		for x := 0; x < len(matrix); x++ {
			if len(matrix[x]) > 0 {
				temp = append(temp, matrix[x])
			}
		}
		matrix = temp

		if direction == "forward" {
			direction = "backward"
		} else {
			direction = "forward"
		}

	}
	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})) // [1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]
}
