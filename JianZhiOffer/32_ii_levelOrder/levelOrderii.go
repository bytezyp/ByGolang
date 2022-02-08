package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	stack := [][]*TreeNode{{root}}
	for len(stack) > 0 {
		node := stack[0]
		arr := []int{}
		for _, v := range node {
			arr = append(arr, v.Val)
		}
		ret = append(ret, arr)
		brr := []*TreeNode{}
		for _, v := range node {
			if v.Left != nil {
				brr = append(brr, v.Left)
			}
			if v.Right != nil {
				brr = append(brr, v.Right)
			}
		}
		if len(brr) != 0 {
			stack = append(stack, brr)
		}
		stack = stack[1:]

	}
	return ret
}
func main() {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	ret := levelOrder(tree)
	fmt.Println(ret)
}
