package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	level := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var curLevel []int
		for i := 0; i < size; i++ {
			// 循环去拿出，当前层的节点，size 是层的节点数量，其实queue，只保留下一层的节点数量
			node := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, node.Val)
			// 这里注意 node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 这里注意 node.Right
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 用于判断层的奇偶层
		if level%2 == 1 {
			ret = append(ret, curLevel)
		} else {
			ret = append(ret, reverseArray(curLevel))
		}
		level++
	}
	return ret
}
func reverseArray(array []int) []int {
	n := len(array)
	// 除以2 只需要交换4次，如果除以len，就交换回去了
	for i := 0; i < n/2; i++ {
		array[i], array[n-1-i] = array[n-1-i], array[i]
	}

	return array
}
func main() {
	tree := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}}}
	ret := levelOrder(tree)
	//arr := []int{1,2,3,4}
	//fmt.Println(reverseArray(arr))
	fmt.Println(ret)
}
