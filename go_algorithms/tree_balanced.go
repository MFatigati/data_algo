/*
Base case: count = 1
*/

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(left int, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	rightDepth := maxDepth(root.Left)
	leftDepth := maxDepth(root.Right)

	difference := math.Abs(float64(rightDepth) - float64(leftDepth))
	return difference <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {

}
