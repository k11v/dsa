package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// LeetCode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultPrev := &ListNode{Val: 0, Next: nil}

	a := l1
	b := l2
	c := resultPrev
	carry := 0

	for a != nil || b != nil || carry != 0 {
		x := carry
		if a != nil {
			x += a.Val
			a = a.Next
		}
		if b != nil {
			x += b.Val
			b = b.Next
		}

		c.Next = &ListNode{Val: x % 10, Next: nil}
		c = c.Next
		carry = x / 10
	}

	result := resultPrev.Next
	return result
}
