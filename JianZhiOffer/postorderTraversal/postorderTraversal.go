package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	var pre *TreeNode
	ret := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			ret = append(ret, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}

	}
	return ret
}
func main() {
	tree := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}}}
	ret := postOrderTraversal(tree)
	fmt.Println(ret)
}
