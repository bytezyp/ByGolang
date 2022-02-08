package main

import "fmt"

func appendArr(arr []int) {
	arr = append(arr, 4)
	//arr[2] = 5

}

func main() {
	arr := []int{1, 2, 3}
	arr = make([]int, 0)
	appendArr(arr)
	fmt.Println(arr)
	a := new(int) // 返回T类型的新分配内存控制键的，返回指针
	fmt.Printf("%v, %p, %T, %v", a, a, a, *a)
	// 输出 ： 0xc00001a0b0, 0xc00001a0b0, *int, 0
	// make() 返回初始化之后的T类型的值

	//arr1 := new([]int)
	arr1 := make([]int, 0)
	// append([]type, elems...) 第一个元素添加切片，不是指针，所以不能append,append 返回的是一个新的地址
	arr1 = append(arr1, 1)

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	//arr1 := [4]int{1, 2, 3, 4}
	//slice := arr1[:2]
	//fmt.Println(slice)
	//for i := 1; i <= 8; i++ {
	//	slice = append(slice, 1)
	//}
	//fmt.Println(slice, arr1)
	//arr := []int{1,2,3,4,5}
	//arr := []int{3,2,4}
	//fmt.Println(twoSum(arr, 6))

}
func twoSum(nums []int, target int) []int {
	arr := make([]int, 0)
	for i := 0; i <= len(nums); i++ {
		n := i + 1
		for ; n < len(nums); n++ {
			sum := nums[i] + nums[n]
			if sum == target {
				new := []int{i, n}
				arr = append(arr, new...)
			}
		}
	}
	return arr
}
