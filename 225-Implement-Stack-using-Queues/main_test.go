package main

// LeetCode

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}


func (this *MyStack) Push(x int)  {
	this.q = append(this.q, x)
}


func (this *MyStack) Pop() int {
	if len(this.q) == 0 {
		panic("stack is empty")
	}
	var x int
	for i := 0; i < len(this.q) - 1; i++ {
		x, this.q = dequeue(this.q)
		this.q = append(this.q, x)
	}
	x, this.q = dequeue(this.q)
	return x
}


func (this *MyStack) Top() int {
	if len(this.q) == 0 {
		panic("stack is empty")
	}
	var x int
	for i := 0; i < len(this.q) - 1; i++ {
		x, this.q = dequeue(this.q)
		this.q = append(this.q, x)
	}
	x, this.q = dequeue(this.q)
	this.q = append(this.q, x)
	return x
}


func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

func dequeue(q []int) (int, []int) {
	return q[0], q[1:]
}
