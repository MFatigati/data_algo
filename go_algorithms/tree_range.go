package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val >= low && root.Val <= high {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	} else {
		return rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
}

func main() {

	node1 := TreeNode{}
	node1.Val = 10
	node2 := TreeNode{}
	node2.Val = 5
	node3 := TreeNode{}
	node3.Val = 3
	node4 := TreeNode{}
	node4.Val = 7
	node5 := TreeNode{}
	node5.Val = 15
	node6 := TreeNode{}
	node6.Val = 18
	node1.Left = &node2
	node2.Left = &node3
	node2.Right = &node4
	node1.Right = &node5
	node5.Right = &node6

	fmt.Println(rangeSumBST(&node1, 7, 15))

}
