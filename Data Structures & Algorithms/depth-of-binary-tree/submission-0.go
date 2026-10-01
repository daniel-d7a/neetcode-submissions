/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    return calcDepth(root)

}

func calcDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	right := calcDepth(node.Right)
	left := calcDepth(node.Left)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

