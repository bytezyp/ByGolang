package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法 一
//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
//你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {

		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2 // 赋值后面的节点
		// 前两个节点，位置颠倒
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

// 方法 二
// 利用栈 两个长度的
func swapPairs2(head *ListNode) *ListNode {
	stack := make([]*ListNode, 2, 2)
	dump := &ListNode{0, nil}
	p := dump
	cur := head
	for cur != nil && cur.Next != nil {
		node1 := cur
		node2 := cur.Next
		cur = node2.Next
		stack[1] = node2
		stack[0] = node1
		dump.Next = stack[1]
		dump = dump.Next
		dump.Next = stack[0]
		stack = make([]*ListNode, 2, 2)
		dump = dump.Next
	}
	if cur != nil {
		dump.Next = cur
	}
	return p.Next
}
func main() {
	stack := make([]*ListNode, 0, 2)
	fmt.Println(stack, len(stack), cap(stack))
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	//head.Next.Next.Next.Next = &ListNode{5,nil}
	//fmt.Println(head, head.Next,head.Next.Next )
	pp := swapPairs2(head)
	fmt.Println(pp, pp.Next, pp.Next.Next, pp.Next.Next.Next, pp.Next.Next.Next.Next)
	return
	swapPairs(head)
}
