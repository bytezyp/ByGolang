package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main1() {
	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)
	//s4 := make([]string,0) 和切片类型无关
	if s1 == nil {
		fmt.Println("s1 silce is nil")
	}
	fmt.Println(s1)
	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%v,\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
