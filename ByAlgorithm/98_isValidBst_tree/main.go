package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 验证是否是二叉搜索树 左子树 < root < 右子树 ，左右子树，也是
// 中序遍历 用栈和循环的方式
func isValidBst(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt
	for root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 最左边的叶子节点的值，默认肯定大于math.minInt
		if inorder >= root.Val {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func best(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true

}

func main() {
	fmt.Println(math.MinInt)
}
