package main

import "container/list"

// 用栈定义队列
type CQueue struct {
	stack1, stack2 *list.List
}

// 初始化队列
func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		ele := this.stack2.Back()
		this.stack2.Remove(ele)
		return ele.Value.(int) // 取出元素，断言自己放入的类型
	}
	return -1
}
