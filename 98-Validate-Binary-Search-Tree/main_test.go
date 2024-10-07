package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// LeetCode

func isValidBST(root *TreeNode) bool {
	var isValidBSTWithLeftAndRight func(root *TreeNode, leftVal *int, rightVal *int) bool
	isValidBSTWithLeftAndRight = func(root *TreeNode, leftVal *int, rightVal *int) bool {
		if root == nil {
			return true
		}
		if !((leftVal == nil || root.Val > *leftVal) && (rightVal == nil || root.Val < *rightVal)) {
			return false
		}
		return isValidBSTWithLeftAndRight(root.Left, leftVal, &root.Val) && isValidBSTWithLeftAndRight(root.Right, &root.Val, rightVal)
	}
	return isValidBSTWithLeftAndRight(root, nil, nil)
}
