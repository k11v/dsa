package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// LeetCode

func reverseList(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	curr := head
	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}
