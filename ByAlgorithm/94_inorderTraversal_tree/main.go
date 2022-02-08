package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树 中序遍历(左子树 中 右子树)的方式循环方式
func inOrderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	// root 判断是否到尽头，stack 用于判断往回查找出站
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1] // 左闭区间，右开
		root = root.Right
	}
	return
}

// 二叉树 中序遍历(左中右)的 递归方式
func inOrderTraversalV2(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	// 运用变量作用域，res 是全局的
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func main() {

}
