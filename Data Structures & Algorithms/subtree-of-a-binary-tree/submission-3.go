/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    
	if root == nil {
		return false
	}

	if dfs(root, subRoot) {
		return true
	}

	left := isSubtree(root.Left, subRoot)
	right := isSubtree(root.Right, subRoot)
	
	return left || right
}

func dfs(first, second *TreeNode) bool {
	if first == nil && second == nil{
		return true
	} else if first == nil || second == nil {
		return false
	}

	right := dfs(first.Right, second.Right)
	left := dfs(first.Left, second.Left)

	same := first.Val == second.Val

	return same && left && right
}