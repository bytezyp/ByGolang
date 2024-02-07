package main

import "fmt"

func sortCount(nums []int)[]int  {
	n := len(nums)
	if n <2 {
		return nums
	}
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] >max {
			max = nums[i]
		}
	}
	count := make([]int, max+1)
	// 计数
	for i := 0; i < n; i++ {
		count[nums[i]]++
	}
	output := make([]int,0,n)
	for k, v := range count {
		for; v >0;v-- {
			output = append(output, k)
		}
	}
	return output
}
/**
计数排序是一种通过计数而不是比较进行排序的算法，适用于范围较小的整数序列。
它的优势在于在对一定范围内的整数排序时，时间复杂度为Ο(n+k)（其中k是整数的范围），快于任何比较排序算法。
当然这是一种牺牲空间换取时间的做法。
 */
func main()  {
	arr := []int{1,3,77,44,4,33}
	fmt.Print(arr)
	fmt.Print(sortCount(arr))

}
