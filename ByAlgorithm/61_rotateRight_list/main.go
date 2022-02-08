package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	//判断参数
	if head == nil || k <= 0 || head.Next == nil {
		return head
	}
	iter := head
	// 求链表长度
	n := 1
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	// 求模 返回头指针
	add := n - k%n
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
func main() {

}
