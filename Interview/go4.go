package main

import "fmt"

var (
	size     = 1024
	max_size = size * 2
)

func main() {
	new1()
	slice1()
}
func new1() {
	//list := new([]int) 编译不过
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
}
func slice1() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// s1 = append(s1, s2) 错误
	s1 = append(s1, s2...) // 正确
	fmt.Println(s1)
}
