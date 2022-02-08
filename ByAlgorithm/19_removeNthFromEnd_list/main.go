package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getListLen(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEndV2(head *ListNode, num int) *ListNode {
	nodes := []*ListNode{}
	dump := &ListNode{0, head}
	for node := dump; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	// 减去1 ，因为之前，加了一个节点
	prev := nodes[len(nodes)-num-1]
	prev.Next = prev.Next.Next
	return dump.Next
}

// 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, num int) *ListNode {
	length := getListLen(head)
	// 前面加个节点用于保存头节点用
	dump := &ListNode{0, head}
	cur := dump
	for i := 0; i < length-num; i++ {
		cur.Next = cur.Next.Next
	}
	return dump.Next
}
func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	removeNthFromEnd(head, 3)
	fmt.Println(head)
}
