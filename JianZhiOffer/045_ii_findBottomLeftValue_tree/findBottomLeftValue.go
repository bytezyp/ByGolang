package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var ret int
	if root == nil {
		return ret
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		size := len(stack)
		ret = stack[0].Val
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:len(stack)]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return ret
}

func main() {
	tree := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}}}
	ret := findBottomLeftValue(tree)
	fmt.Println(ret)
}
