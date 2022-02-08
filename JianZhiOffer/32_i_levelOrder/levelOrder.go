package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层级遍历
func levelOrder(root *TreeNode) []int {
	stack := []*TreeNode{root}
	ret := []int{}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		ret = append(ret, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}
	return ret
}

func main() {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	ret := levelOrder(tree)
	fmt.Println(ret)
}
