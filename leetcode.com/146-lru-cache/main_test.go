package main

import (
	"testing"
	"reflect"
)

func TestLRUCache(t *testing.T) {
	got := make([]int, 0)
	want := make([]int, 0)

	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	got, want = append(got, c.Get(1)), append(want, 1)
	c.Put(3, 3)
	got, want = append(got, c.Get(2)), append(want, -1)
	c.Put(4, 4)
	got, want = append(got, c.Get(1)), append(want, -1)
	got, want = append(got, c.Get(3)), append(want, 3)
	got, want = append(got, c.Get(4)), append(want, 4)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// LeetCode

type node struct {
	key int
	value int
	next *node
	prev *node
}

type LRUCache struct {
	capacity int
	inner map[int]*node
	head *node
	tail *node
}

func Constructor(capacity int) LRUCache {
	if capacity < 0 {
		panic("capacity is negative")
	}
	head := &node{}
	tail := &node{}
	head.prev = tail
	tail.next = head
	return LRUCache{capacity: capacity, inner: make(map[int]*node), head: head, tail: tail}
}

func (c *LRUCache) Get(key int) int {
	n, present := c.inner[key]
	if !present {
		return -1
	}
	unlink(n)
	linkAtBack(c.head, n)
	return n.value
}

func (c *LRUCache) Put(key int, value int)  {
	if n, present := c.inner[key]; present {
		unlink(n)
		linkAtBack(c.head, n)
		n.value = value
		return
	}

	if len(c.inner) == c.capacity {
		if c.tail.next == c.head {
			return
		}
		delete(c.inner, c.tail.next.key)
		unlink(c.tail.next)
	}

	n := &node{key: key, value: value}
	c.inner[key] = n
	linkAtBack(c.head, n)
}

func unlink(n *node) {
	n.next.prev = n.prev
	n.prev.next = n.next
	n.next = nil
	n.prev = nil
}

func linkAtBack(nodeWithBack *node, n *node) {
	n.next = nodeWithBack
	n.prev = nodeWithBack.prev
	nodeWithBack.prev.next = n
	nodeWithBack.prev = n
}
