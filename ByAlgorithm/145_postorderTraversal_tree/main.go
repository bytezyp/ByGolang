package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树 后续遍历  迭代方式，左、右、中
func postorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			res = append(res, root.Value)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}

// 递归方式
func postorderTraversalV2(root *TreeNode) (res []int) {
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node != nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Value)
	}
	postorder(root)
	return res
}

func main() {

}
