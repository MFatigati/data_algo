/*

 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeBFS(root *TreeNode) {
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		if popped.Left != nil {
			queue = append(queue, popped.Left)
		}
		if popped.Right != nil {
			queue = append(queue, popped.Right)
		}
		fmt.Println(popped)
	}
}

func main() {

	node1 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 0}
	node6 := &TreeNode{Val: 7}
	node7 := &TreeNode{Val: 8}
	node1.Left = node2
	node2.Left = node3
	node2.Right = node4
	node1.Right = node5
	node5.Left = node6
	node5.Right = node7

	treeBFS(node1)

}
