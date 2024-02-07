package main

import (
	"fmt"
	"time"
)

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
	//outerSlice := make([]string, 0, 4)
	//outerSlice = append(outerSlice, "a", "a")
	//fmt.Printf("%p %v \n", &outerSlice, cap(outerSlice))
	//modifySlice(outerSlice)
	//fmt.Println(outerSlice)

	str := "!abbc! 123q中文"
	strs := []rune(str)
	newStr := make([]rune, 0, 100)
	for i := 0; i < len(strs); i++ {
		if (97 <= strs[i] && strs[i] <= 122) || strs[i] == 47 || strs[i] == 38 || strs[i] == 39 || strs[i] == 32 || strs[i] == 44 {
			newStr = append(newStr, strs[i])
		}
	}
	fmt.Println(string(newStr))

	//var num *int // 定义int指针
	//num2 := 22
	//num = &num2 // 取num2中的地址，复制给指针  (指针类型值，只能复制地址，& 代表取其他类型值的地址)
	////fmt.Println(num)
	//fmt.Println(*num)
	adxBlackVal := time.Now().AddDate(0, 0, 0).Format("2006010205")
	now := time.Now().In(time.Local)
	fmt.Println(adxBlackVal, now)
	//t := time.Duration(1234) * time.Second
	//t2 := 123 * time.Second
	var arr []int
	println(len(arr), 11111)

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
