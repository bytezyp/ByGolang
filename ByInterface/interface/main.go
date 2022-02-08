package main

import "fmt"

// 数组切片，不能转化 接口切片 最主要原因是底层结构不一样
func sliceSum(inters []interface{}) (res interface{}) {
	//nums := inters.([]int)
	sum := 0
	for _, inter := range inters {
		sum += inter.(int)
	}
	return sum
}

func main() {
	is := []int{1, 2, 3}
	iis := make([]interface{}, len(is))
	for i, i2 := range is {
		iis[i] = i2
	}
	fmt.Println(sliceSum(iis))
}
