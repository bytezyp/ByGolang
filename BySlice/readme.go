package main

import (
	"fmt"
	"time"
)

func modifySlice(innerSlice []string) {
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice, len(innerSlice),cap(innerSlice))
}
func main() {
	outerSlice:= make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	modifySlice(outerSlice)
	fmt.Println(outerSlice,len(outerSlice),cap(outerSlice))
	return
	var ints []int
	fmt.Printf("ints : %v , %p \n", ints, ints)
	// ints : [] , 0x0
	ints3 := make([]int, 0, 3)
	fmt.Printf("ints3 : %v, %v, %v \n", ints3, len(ints3), cap(ints3))
	// ints3 : [], 0, 3  len:0, cap: 3
	ints2 := make([]int, 2, 5)
	fmt.Printf("ints2 : %v , %p len:%v cap:%v \n", ints2, ints2, len(ints2), cap(ints2))
	// ints2 : [0 0] , 0xc0000ba000 len:2 cap:5
	ints2 = append(ints2, 2)
	fmt.Printf("ints2 : %v , %p len:%v cap:%v \n", ints2, ints2, len(ints2), cap(ints2))
	// ints2 : [0 0 2] , 0xc0000ba000 len:3 cap:5
	// ints2[3] 超过 len()的长度，就会出现越界访问，报panic错误
	// fmt.Println(ints2[3])

	//字符串类型
	ps := new([]string)
	fmt.Printf("ps : %v , %p \n", ps, ps)
	// ps : &[] , 0xc00000c0e0
	// 分配 内存
	*ps = append(*ps, "eggo")
	fmt.Printf("ps : %v , %p \n", ps, ps)
	fmt.Printf("ps : %v , %p len:%v cap:%v \n", ints2, ints2, len(*ps), cap(*ps))

	arr := make([]int, 3, 3)
	for i := 0; i < 3; i++ {
		//fmt.Println(arr[i])
		go func(i int, arr []int) {
			if i%2 == 0 {
				arr[i] = i + 1
			} else {
				arr[i] = i + 1
			}
			fmt.Println(i)
		}(i, arr)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(arr)
}
