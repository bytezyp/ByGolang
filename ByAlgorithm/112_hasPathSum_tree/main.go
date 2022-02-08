package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	que := []*TreeNode{root}
	queVal := []int{root.Val}
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		val := queVal[0]
		queVal = queVal[1:]
		if node.Left != nil && node.Right != nil {
			if val == sum {
				return true
			}
			continue
		}
		if node.Left != nil {
			que = append(que, node.Left)
			// node.Left.val
			queVal = append(queVal, val+node.Left.Val)
		}
		if node.Right != nil {
			que = append(que, node.Right)
			// 注意 是 node.Right.val
			queVal = append(queVal, node.Right.Val+val)
		}
	}
	return false
}

func main() {

}
