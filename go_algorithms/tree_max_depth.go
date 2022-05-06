/*
P
	I:
	O:
E
D
A
*/

package main

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

}
