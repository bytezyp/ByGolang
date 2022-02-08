package main

import "fmt"

const a = 1 << iota
const b = iota

func modifySlice(innerSlice []string) {
	fmt.Printf("%p %v \n", &innerSlice, cap(innerSlice))
	innerSlice = append(innerSlice, "a")
	fmt.Printf("%p %v \n", &innerSlice, cap(innerSlice))
	fmt.Println(innerSlice)
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}

func main() {
	outerSlice := make([]string, 0, 4)
	outerSlice = append(outerSlice, "a", "a")
	fmt.Printf("%p %v \n", &outerSlice, cap(outerSlice))
	modifySlice(outerSlice)
	fmt.Println(outerSlice)
}

//func main() {
//
//	fmt.Println(a, b)
//	str := "hello"
//	//str[0] = 'x' // 这样是错误的，go 语言的字符传是只读的，可以转成字节的方式
//	fmt.Println(str)
//	str2 := []byte(str)
//	fmt.Println(string(str2))
//	str2[0] = 'x'
//	fmt.Println(string(str2))
//}
