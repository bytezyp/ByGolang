package main

import "fmt"

func printAll[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

//func findFunc[T comparable](a []T, v T) T {
//	for i, e := range a {
//		if e == v {
//
//		}
//	}
//	return -1
//}

// 感觉泛式很方便，真的不用断言，有点动态语言的方式
func main() {
	//intSlice := []int{1, 2, 4, 5, 7, 8}
	//printAll(intSlice)
	//strSlice := []string{"aa", "bb", "cc", "ee"}
	//printAll(strSlice)

	//fmt.Println(Add(1, 2))

}
