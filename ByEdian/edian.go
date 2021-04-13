package main

import (
	"fmt"
	"unsafe"
)

const INI_SIZE = int(unsafe.Sizeof(0)) // 64位操作系统， 8bytes
func main() {
	fmt.Println(INI_SIZE)
}
