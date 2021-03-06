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

func flatten1(root *TreeNode) {
	preOrderSlice := preOrderTraversal(root)

	for i, elem := range preOrderSlice {
		if i < len(preOrderSlice)-1 {
			elem.Right = preOrderSlice[i+1]
		}
		elem.Left = nil
	}
}

func preOrderTraversal(current *TreeNode) (result []*TreeNode) {
	if current == nil {
		return
	}

	result = append(result, current)
	result = append(result, preOrderTraversal(current.Left)...)
	result = append(result, preOrderTraversal(current.Right)...)

	return
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
	node6 := TreeNode{}
	node6.Val = 6

	node1.Left = &node2
	node1.Right = &node5
	node2.Left = &node3
	node2.Right = &node4
	node5.Right = &node6

	flatten1(&node1)

}
