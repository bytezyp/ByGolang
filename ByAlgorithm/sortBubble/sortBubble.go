package main

import "fmt"

// 冒泡排序
func sortBubble(nums []int) []int {
	numLen := len(nums)
	if numLen <0 {
		return nums
	}
	for i := 0; i < numLen; i++ {
		for j := 0; j < numLen-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
	return nums
}
// 错误排序
func sortBubbleV2(nums []int) []int {
	numLen := len(nums)
	if numLen <0 {
		return nums
	}
	for i := 0; i < numLen; i++ {
		for j := numLen-1; j <numLen-i ; j-- {
			fmt.Print(nums[j],nums[j-1])
			if nums[j] > nums[j-1] {
				nums[j],nums[j-1] = nums[j-1],nums[j]
			}
		}
	}
	return nums
}
/**
冒泡排序是一种比较简单的排序算法，它的原理是：
从左到右遍历数组，相邻元素两两进行比较。每次比较完一轮，就会找到数组中最大的一个或最小的一个。这个数就会从序列的最右边冒出来。
以从小到大排序为例，第一轮比较后，所有数中最大的那个数就会浮到最右边；第二轮比较后，所有数中第二大的那个数就会浮到倒数第二个位置……就这样一轮一轮地比较，最后实现从小到大排序。
 */
func main()  {
	//arr := []int{5,7,4,3,11}
	//fmt.Print(arr)
	//fmt.Print(sortBubble(arr))
	brr := []int{5,7,4,3,11}
	fmt.Print(sortBubbleV2(brr))
}

