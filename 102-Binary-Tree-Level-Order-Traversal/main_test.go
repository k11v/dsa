package main

import (
	"testing"
	"reflect"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func TestLevelOrder(t *testing.T) {
	tests := []struct{
		root *TreeNode
		want [][]int
	}{
		{
			&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			&TreeNode{1, nil, nil},
			[][]int{{1}},
		},
		{
			nil,
			[][]int{},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := levelOrder(tt.root), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode

type nodeWithLevel struct {
	node *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	r := make([][]int, 0)

	q := make([]nodeWithLevel, 0)
	q = append(q, nodeWithLevel{root, 0})

	for len(q) != 0 {
		var nl nodeWithLevel
		nl, q = unprepend(q)

		if nl.level + 1 != len(r) {
			r = append(r, make([]int, 0))
		}
		r[nl.level] = append(r[nl.level], nl.node.Val)

		if left := nl.node.Left; left != nil {
			q = append(q, nodeWithLevel{left, nl.level + 1})
		}
		if right := nl.node.Right; right != nil {
			q = append(q, nodeWithLevel{right, nl.level + 1})
		}
	}

	return r
}

func unprepend(q []nodeWithLevel) (nodeWithLevel, []nodeWithLevel) {
	return q[0], q[1:]
}
