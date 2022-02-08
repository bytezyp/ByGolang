package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 难懂 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == p.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 存储父节点的方式
func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {
	// 定义 root 的所有父节点
	parent := map[int]*TreeNode{}
	// 记录所有访问过的
	visited := map[int]bool{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}

		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}

	dfs(root)
	// p 向上找到所有p的父节点
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	// p 向上找和q相同的祖先
	for q != nil {
		if visited[p.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func main() {

}
