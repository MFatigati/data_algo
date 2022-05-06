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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == nil && q == nil {
			return true
		} else {
			return false
		}
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// func isSymmetric(root *TreeNode) bool {
// 	left := invertTree(root.Left)
// 	right := root.Right
// 	return isSameTree(left, right)
// }

// solution not based on prior functions below

func isSymmetric(root *TreeNode) bool {
	return symmetricHelper(root.Left, root.Right)
}

func symmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	return left.Val == right.Val
		&& symmetricHelper(left.Left, right.Right)
		&& symmsymmetricHelper(left.Right, right.Left)
}

func main() {

}
