package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var node *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = node
		node = cur
		cur = next
	}
	return node
}

func main() {

}
