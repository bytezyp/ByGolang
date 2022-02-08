package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树，前序遍历，迭代方式， 中、左、右
func preorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Value)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		//  左是闭区间，右是开区间， 【0，1）
		stack = stack[:len(stack)-1]
	}
	return
}

// 递归方式，前序遍历
func preorderTraversalV2(root *TreeNode) (res []int) {
	var preorderFunc func(*TreeNode)
	preorderFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Value)
		preorderFunc(root.Left)
		preorderFunc(root.Right)
	}
	preorderFunc(root)
	return
}
func main() {

}
