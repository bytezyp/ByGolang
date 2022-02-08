package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{7, 8, 9}
	fmt.Printf("slice2 copy before :%v slice1 : %v \n", slice2, slice1)
	copy(slice2, slice1)
	fmt.Printf("slice2 copy after :%v slice1： %v \n", slice2, slice1)
	// 把slice1中前三个元素放在slice2中
	fmt.Println("把slice1中前三个元素放在slice2中")
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{7, 8, 9}
	fmt.Printf("slice3 copy before :%v slice4 : %v \n", slice3, slice4)
	copy(slice3, slice4)
	fmt.Printf("slice3 copy after :%v slice4 : %v \n", slice3, slice4)
	fmt.Println("把slice4中的3个元素，放到slice3的前3个位置中")
}
