/*

 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (acc [][]int) {
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}
		queueLength := len(queue)

		for counter := 0; counter < queueLength; counter++ {
			popped := queue[0]
			queue = queue[1:]

			level = append(level, popped.Val)

			if popped.Left != nil {
				queue = append(queue, popped.Left)
			}
			if popped.Right != nil {
				queue = append(queue, popped.Right)
			}
		}
		acc = append(acc, level)
	}

	return acc
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

	fmt.Println(levelOrder(node1))

}
