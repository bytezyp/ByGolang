package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	//for l != nil {
	//	fmt.Println(l.Val)
	//	l = l.Next
	//}
	pre := reverseList(l)
	for pre != nil {
		fmt.Println(pre.Val)
		pre = pre.Next
	}
}
