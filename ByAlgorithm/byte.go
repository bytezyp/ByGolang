package main

import (
	"fmt"
	"unsafe"
)

type St1 struct {
	A byte
	B int
	C byte
}

func main() {
	fmt.Println("st1 占用的字节数：" + fmt.Sprint(unsafe.Sizeof(St1{}.A)))
	fmt.Println("st1 对齐的字节数：" + fmt.Sprint(unsafe.Alignof(St1{}.B)))
	fmt.Println("ST1结构体 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(St1{})))
	fmt.Println("ST1结构体 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(St1{})))
}
