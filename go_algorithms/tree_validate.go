/*

 */

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	max := math.MaxInt64
	min := math.MinInt64
	return validBSTHelper(root, min, max)
}

/*
	9
   4  15
 2 3  11 30

*/
func validBSTHelper(current *TreeNode, min int, max int) bool {
	if current == nil {
		return true
	}
	isCurrentValid := current.Val < max && current.Val > min

	return isCurrentValid && validBSTHelper(current.Left, min, current.Val) && validBSTHelper(current.Right, current.Val, max)
}

func main() {

	node1 := TreeNode{}
	node1.Val = 1
	node2 := TreeNode{}
	node2.Val = 2
	node3 := TreeNode{}
	node3.Val = 3
	node4 := TreeNode{}
	node4.Val = 4
	node5 := TreeNode{}
	node5.Val = 5
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node2.Right = &node5

	fmt.Println(isValidBST(&node1))
}
