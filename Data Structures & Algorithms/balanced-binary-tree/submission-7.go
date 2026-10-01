/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var balanced = true

func isBalanced(root *TreeNode) bool {
	balanced = true
	if root == nil {
		return true
	}
    
	calcDepth(root)

	fmt.Println("res", balanced)
	return balanced
}

func calcDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	right := calcDepth(node.Right)
	left := calcDepth(node.Left)

	fmt.Println("left: ", left)
	fmt.Println("right: ", right)

	if math.Abs(float64(left - right)) > 1 {
		fmt.Println("changed")
		balanced = false
	}

	return 1 + max(left, right)
}