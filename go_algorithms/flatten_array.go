/*

Count depth
	If at the appropriate depth, return current slice
	Else, return outermost array plus flattened inner arrays
5 layer array, flatten to 1

The optional parameter specifies how many levels to flatten, not up to what flatness

*/

package main

import "fmt"

type mixed interface{}

func flatten(current []interface{}, depth int) []interface{} {
	result := []interface{}{}

	if depth == 0 {
		return current
	} else {
		for _, elem := range current {
			switch t := elem.(type) {
			case int:
				result = append(result, elem)
			case mixed:
				for i := 0; i < ; i++ {
				}
			default:
				fmt.Println(t)
			}
		}
	}

	return result

	//return flatten(current, depth-1)
}

func main() {
	fmt.Println(flatten([]interface{}{1, 2, []interface{}{3, 4, []interface{}{5, 6}}}, 1))
}
