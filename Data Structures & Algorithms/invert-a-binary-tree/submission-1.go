/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    swapChildren(root)
	return root
}

func swapChildren(node *TreeNode) {

	if node == nil {
		return
	}

	oldLeft := node.Left
	oldRight := node.Right

	node.Left = oldRight
	node.Right = oldLeft

	swapChildren(node.Left)
	swapChildren(node.Right)
}