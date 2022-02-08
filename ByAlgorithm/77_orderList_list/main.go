package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 运用快慢指针的方式找到中点，然后用归并排序的方式，进行方式，排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next, slow = nil, slow.Next
	//fmt.Println(slow.Val,slow.Next, slow)
	return merge(sortList(head), sortList(slow))
}

func merge(n1, n2 *ListNode) *ListNode {
	if n1 == nil {
		return n2
	}

	if n2 == nil {
		return n1
	}
	p1 := n1
	p2 := n2
	queue := &ListNode{}
	curr := queue
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	}
	if p2 != nil {
		curr.Next = p2
	}
	return queue.Next
}

func main() {
	l := &ListNode{4, &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{5, nil}}}}}
	new := sortList(l)
	for new != nil {
		fmt.Println(new.Val)
		new = new.Next
	}

	n3 := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(n3.Val, n3.Next)
	n3.Next, n3 = nil, n3.Next
	fmt.Println(n3.Val, n3.Next)

}
