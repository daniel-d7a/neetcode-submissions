/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    return dfs(p, q)
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