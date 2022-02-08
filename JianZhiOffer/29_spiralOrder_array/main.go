package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		// 定义行，列
		rows, columns = len(matrix), len(matrix[0])
		// 初始化 元素的切片
		order = make([]int, rows*columns)
		// 定义起始
		index = 0
		// 定义左 右 上 下
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	// 循环条件
	for left <= right && top <= bottom {
		// 列做运算，从左至右
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		// 行做运算，从上到下
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			// 从右向左
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			// 从底向上
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ret := spiralOrder(arr)
	fmt.Println(ret)
}
