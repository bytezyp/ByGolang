package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	arr := []int{}
	n := 0
	cur := head
	for cur != nil {
		arr = append(arr, cur.Value)
		cur = cur.Next
		if cur != nil {
			n++
		}
	}
	brr := []int{}
	for n >= 0 {
		brr = append(brr, arr[n])
		n--
	}
	return brr
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(reversePrint(l))
}
