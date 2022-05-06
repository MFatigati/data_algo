package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumHelper(root, 0)
}

func sumHelper(currentNode *TreeNode, sum int) int {
	if currentNode == nil {
		return sum
	}
	sum = sum*10 + currentNode.Val

	if currentNode.Left != nil && currentNode.Right != nil {
		sum = sumHelper(currentNode.Left, sum) + sumHelper(currentNode.Right, sum)
	} else {
		if currentNode.Right != nil {
			sum = sumHelper(currentNode.Right, sum)
		} else {
			sum = sumHelper(currentNode.Left, sum)
		}
	}

	fmt.Println(sum)

	return sum
}

func main() {

	node1 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 0}
	node1.Left = node2
	node2.Left = node3
	node2.Right = node4
	node1.Right = node5

	fmt.Println(sumNumbers(node1))

}
