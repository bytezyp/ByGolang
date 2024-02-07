package main

import "fmt"

func sortSelect(nums []int)[]int  {
	numsLen := len(nums)
	if numsLen <0 {
		return nums
	}
	for i := 0; i < numsLen; i++ {
		min := i
		for j := i+1; j < numsLen; j++ {
			if nums[j] < nums[min]{
				nums[min],nums[j] = nums[j],nums[min]
			}
		}
		nums[i] = nums[min]
	}
	return nums
}
/**
选择排序的原理是，首先在开始的时候遍历整个数组，找到序列中的最小元素，然后将这个元素与第一个元素交换，这样最小的元素就放到了它的最终位置上了。
然后，从第二个元素开始遍历，找到剩下的n-1个元素中的最小元素，再与第二个元素进行交换。
以此类推，直到最后一个元素放到了最终位置。
 */
func main()  {
	brr := []int{5,7,4,3,22,0,11}
	fmt.Print(sortSelect(brr))
}
