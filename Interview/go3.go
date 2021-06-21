package main

import "fmt"

func main3() {
	make1() // 会有初始值
	make2() // 不会有初始值
}
func make1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
func make2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 返回形参，如果是多个必选用括号整理好，如果一个有形参名称，必要都要做有，同时也要有括号。
func funcMui(x, y int) (sum, aa int) {
	return x + y, 1
}
