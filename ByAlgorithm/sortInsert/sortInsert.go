package main

import "fmt"

func sortInsert(nums []int) []int {
	numsLen := len(nums)
	if numsLen <0 {
		return nums
	}
	for i := 1; i < numsLen; i++ {
		preIndex := i-1
		nowNums := nums[i]
		for nums[preIndex] > nowNums {
			// 换位置
			nums[preIndex+1] = nums[preIndex]
			if preIndex == 0 {
				preIndex--
				println(preIndex)
				break
			}
			preIndex--
		}
		nums[preIndex+1] = nowNums
	}
	return nums
}
/**
插入排序，也叫直接插入排序。
它的原理是，将一个记录插入到已经排好序的有序表中，从而一个新的、记录数增1的有序表。
在其实现过程使用双层循环，外层循环对除了第一个元素之外的所有元素，内层循环对当前元素前面的有序表进行待插入位置查找，并进行移动。
 */
func main()  {
	arr := []int{5,7,4,3,11}
	fmt.Print(arr)
	fmt.Print(sortInsert(arr))
}

