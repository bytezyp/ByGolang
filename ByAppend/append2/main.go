package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

// append 注意事项 ： https://mp.weixin.qq.com/s/VIssdutZpWhs1auMQlMJnw
func init() {
	var slice0 []int
	// var 初始化切片  切片长度和容量 默认是相等的
	fmt.Printf("var 初始化切片：%v len: %v cap: %v \n", slice0, len(slice0), cap(slice0))
	// 如果指向底层数组的指针被覆盖或者修改（copy、重分配、append触发扩容），
	// 此时函数内部对数据的修改将不再影响到外部的切片，代表长度的len和容量cap也均不会被修改
}
func main() {
	arr := make([]int, 0, 2)
	fmt.Println(arr)
	arr = append(arr, 1)
	fmt.Println(arr)
	var numMax int64
	fmt.Println(numMax)
	return

	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 2
	change(slice...)
	fmt.Println(slice, len(slice), cap(slice), slice[:len(slice)], slice[:len(slice)-1])
	slice1 := slice[0:2]
	change(slice1...)

	fmt.Println(slice, len(slice), cap(slice))
}
