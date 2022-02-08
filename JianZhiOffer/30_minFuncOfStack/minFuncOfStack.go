package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func (this *MinStack) Push(value int) {
	this.stack = append(this.stack, value)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(top, value))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}
