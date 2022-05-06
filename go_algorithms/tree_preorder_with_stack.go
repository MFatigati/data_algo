/*

 */

package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		result = append(result, popped.Val)
		stack = stack[:len(stack)-1]
		if popped.Right != nil {
			stack = append(stack, popped.Right)
		}
		if popped.Left != nil {
			stack = append(stack, popped.Left)
		}
	}
	return result
}

func main() {

}
