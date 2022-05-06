/*
P
	I:
	O:
E
D
A
*/

package main

func preorderTraversal(root *TreeNode) []int {

	return preorderHelper(root)
}

func preorderHelper(node *TreeNode) (result []int) {
	if node == nil {
		return []int{}
	}

	result = append(result, node.Val)
	result = append(result, preorderHelper(node.Left)...)
	result = append(result, preorderHelper(node.Right)...)

	return result
}

func main() {

}
