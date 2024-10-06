package main

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {}

// LeetCode

type NestedIterator struct {
	nestedIntegers []*NestedInteger
	nextNestedIterator *NestedIterator
	next int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{nestedIntegers: nestedList, nextNestedIterator: nil, next: 0}
}

func (ni *NestedIterator) Next() int {
	if !ni.HasNext() {
		panic("nested iterator doesn't have next")
	}

	ni.seekNextInteger()

	if ni.nextNestedIterator != nil {
		x := ni.nextNestedIterator.Next()
		if !ni.nextNestedIterator.HasNext() {
			ni.nextNestedIterator = nil
			ni.next++
		}
		return x
	}

	x := ni.nestedIntegers[ni.next].GetInteger()
	ni.next++
	return x
}

func (ni *NestedIterator) HasNext() bool {
	ni.seekNextInteger()
	return ni.next < len(ni.nestedIntegers)
}

func (ni *NestedIterator) seekNextInteger() {
	for ni.next < len(ni.nestedIntegers) {
		if ni.nestedIntegers[ni.next].IsInteger() {
			return
		}

		if ni.nextNestedIterator != nil {
			return
		}

		nextNestedIterator := Constructor(ni.nestedIntegers[ni.next].GetList())
		if nextNestedIterator.HasNext() {
			ni.nextNestedIterator = nextNestedIterator
			return
		}

		ni.next++
	}
}
