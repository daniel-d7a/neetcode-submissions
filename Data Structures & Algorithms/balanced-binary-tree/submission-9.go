/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
    
	_, balanced := calcDepth(root)
	return balanced
}

func calcDepth(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	right, rightBalanced := calcDepth(node.Right)
	left, leftBalanced := calcDepth(node.Left)

	thisBalanced := math.Abs(float64(left - right)) <= 1

	return 1 + max(left, right), thisBalanced && leftBalanced && rightBalanced
}