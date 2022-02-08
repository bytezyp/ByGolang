package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return &ListNode{}
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 为了保存头节点
	dump := &ListNode{0, nil}
	cur := dump
	// 原来两个链表，都是升序
	// 新链表是通过(拼接)给定的两个链表的所有节点组成
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list2
			list1 = list1.Next
		}
	}
	// 最后肯定剩一节点为nil
	if list2 == nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dump.Next
}
func main() {

}
