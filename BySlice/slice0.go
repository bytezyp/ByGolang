package main

import "fmt"

func main0() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:5]
	s2 := s1[2:4]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Printf("切片首个元素地址: %p \n", &arr[0])
	fmt.Printf("切片第二个元素地址: %p \n", &arr[1])
	fmt.Printf("切片第三个元素地址: %p \n", &arr[2])
	fmt.Printf("切片第四个元素地址: %p \n", &arr[3])
	fmt.Printf("切片首地址：%p \n", arr)
	fmt.Printf("由切片生成的新切片s1：%p \n", s1)
	fmt.Println("切片s1：", s1)
	fmt.Printf("由切片生成的新切片s2：%p \n", s2)
	fmt.Println("切片s2:", s2)
	s3 := append(arr, 11, 22, 33)
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Printf("%p \n", s3)
	fmt.Printf("%p \n", arr)
	fmt.Println(s2, len(s2), cap(s2))
}
