package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LeetCode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	s := make([]*TreeNode, 0)
	s = append(s, p, q)

	for len(s) != 0 {
		var u, v *TreeNode
		u, v, s = s[len(s)-2], s[len(s)-1], s[:len(s)-2]
		if u == nil && v == nil {
			continue
		}
		if u != nil && v != nil && u.Val == v.Val {
			s = append(s, u.Left, v.Left)
			s = append(s, u.Right, v.Right)
			continue
		}
		return false
	}

	return true
}
