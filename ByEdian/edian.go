package main

import (
	"fmt"
	"unsafe"
)

const INI_SIZE = int(unsafe.Sizeof(0)) // 64位操作系统， 8bytes
func main() {
	//fmt.Println(INI_SIZE)
	var i = 0x01020304
	fmt.Println("&i", &i)
	bs := (*[INI_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0x04 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}
}
