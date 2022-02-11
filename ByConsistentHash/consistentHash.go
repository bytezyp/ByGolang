package main

import (
	"fmt"
)

type Mutx struct {
	Name string
}

var mutx Mutx
var a int
var slice []int
var Map map[string]int
func main()  {
	m := make(map[string]interface{}, 10)
	fmt.Println(len(m))
	m["11"] = 11
	fmt.Println(len(m))
	// map 只能 求长度，不能求容量
	s := make([]int, 1, 5)
	fmt.Println(len(s))
	// slice 长度为0，只能append添加元素，不能赋值方式,例如 : s[0] = 1
	// slice 既可以求长度，又可以求容量
	s[0] = 1
	fmt.Println(len(s),cap(s))
	slice = make([]int, 1, 10)
	slice = append(slice, 11)
	Map = make(map[string]int)
	Map["11"] = 11
	//fmt.Printf("%v %T %p \n", mutx, &mutx, &mutx)
	//fmt.Printf("%v %T %p \n", a, &a, &a)
	fmt.Printf("%v %T %p \n", slice, &slice, &slice[0])
	//fmt.Printf("%v %T %p \n", s, &s, &s)
	//fmt.Println(unsafe.Sizeof(slice))
	fmt.Println(cap(slice),len(slice))
	slice = append(slice, 11,22)
	fmt.Printf("%v %T %p \n", slice, &slice, &slice[0])
	fmt.Println(cap(slice),len(slice))
}
