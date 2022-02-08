package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 球链表的最大回文
// 递归方式
func isPalindrome(head *ListNode) bool {
	return _isPalindrome(head, head)
}

func _isPalindrome(head1, head2 *ListNode) bool {
	if head1 != nil {
		result := _isPalindrome(head1.Next, head2) && head1.Val == head2.Val
		if head2 != nil && head2.Next != nil {
			*head2 = *head2.Next
		}
		return result
	}
	return true
}

// 链表反转
func reverseList(node *ListNode) *ListNode {
	var pre *ListNode
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 快慢指针方式
func isPalindromeV2() {

}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	isPalindrome(l)
}
