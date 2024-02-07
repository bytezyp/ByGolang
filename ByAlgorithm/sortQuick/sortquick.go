package main

import "fmt"

func sortQuick(nums []int)[]int  {
	numsLen := len(nums)
	if numsLen <2 {
		return nums
	}
	pivot := nums[0]
	var low,mid,high []int
	for i := 0; i < numsLen; i++ {
		if nums[i] < pivot{
			low = append(low, nums[i])
		}else if nums[i]  == pivot{
			mid = append(mid, nums[i])
		}else {
			high = append(high, nums[i])
		}
	}
	low,high = sortQuick(low), sortQuick(high)
	res := append(append(low,mid...),high...)
	return res
}
/**
快速排序要求我们选择一个基准，根据这个基准为原数组分组，比基准大的一组，比基准小的一组，再重复递归地进行快速排序，重新合并每组排序后的序列，即为排好序的序列。
快速排序需要我们先理解递归的思想
 */
func main()  {
	arr := []int{5,7,4,3,11}
	fmt.Print(arr)
	fmt.Print(sortQuick(arr))
}




