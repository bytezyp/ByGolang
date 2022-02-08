package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, num int) *ListNode {
	hair := &ListNode{0, head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < num; i++ {
			tail = tail.Next
			if tail.Next == nil {
				return hair.Next
			}
		}
		// 定义下一个子循环的起点
		nex := tail.Next
		// 反转指定长度子链表
		head, tail = myReverse(head, tail)
		// 重新设置新的起始点
		pre.Next = head
		// 拼接返回后的节点
		tail.Next = nex
		// 为子循环设置起点
		pre = tail
		// 移动头指针
		head = tail.Next

	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		// 起始节点的next
		nex := p.Next
		// 起始接的next 等于 tail.next
		p.Next = prev
		// prev 向前一位
		prev = p
		// p 等于next
		p = nex
	}
	return tail, head
}

func main() {

}
