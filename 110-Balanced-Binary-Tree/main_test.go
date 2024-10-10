package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LeetCode

func isBalanced(root *TreeNode) bool {
	var isBalancedWithHeight func(root *TreeNode) (bool, int)
	isBalancedWithHeight = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}
		leftBalanced, leftHeight := isBalancedWithHeight(root.Left)
		rightBalanced, rightHeight := isBalancedWithHeight(root.Right)
		balanced := leftBalanced && rightBalanced && abs(leftHeight-rightHeight) <= 1
		height := max(leftHeight, rightHeight) + 1
		return balanced, height
	}
	balanced, _ := isBalancedWithHeight(root)
	return balanced
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
