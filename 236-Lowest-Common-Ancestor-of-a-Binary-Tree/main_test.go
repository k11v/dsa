package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LeetCode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathToP, err := path(root, p)
	if err != nil {
		panic(err)
	}

	pathToQ, err := path(root, q)
	if err != nil {
		panic(err)
	}
	
	i := 0
	for i+1 < len(pathToP) && i+1 < len(pathToQ) && pathToP[i+1] == pathToQ[i+1] {
		i++
	}

	return pathToP[i]
}

func path(root, n *TreeNode) ([]*TreeNode, error) {
	if root == nil {
		return nil, fmt.Errorf("root is nil")
	}

	p := []*TreeNode{root}
	if root.Val == n.Val {
		return p, nil
	}
	if subpath, err := path(root.Left, n); err != nil {
		p = append(p, subpath...)
		return p, nil
	}
	if subpath, err := path(root.Right, n); err != nil {
		p = append(p, subpath...)
		return p, nil
	}
	return nil, fmt.Errorf("node %d is not a descendant of root %d", root.Val, n.Val)
}
