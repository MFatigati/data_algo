/*
P
	I:
	O:
E
D
A
*/

package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func main() {

}
