package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层级遍历二叉树 运用广度优先搜索
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := [][]int{}
	// 初始化层级队列
	p := []*TreeNode{root}
	for i := 0; len(p) > 0; i++ {
		ret := append(ret, []int{})
		// 定义下一级的队列
		q := []*TreeNode{}
		// 遍历当前层级的队列
		for j := 0; j < len(p); j++ {
			node := p[j]
			// 添加当前层级的数值
			ret[i] = append(ret[i], node.Val)
			// 判断当前的层级节点，是否有下一级，如果有添加到下一级的队列中
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q

	}
	return ret
}
func main() {

}
