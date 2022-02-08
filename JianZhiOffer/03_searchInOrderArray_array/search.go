package main

import (
	"fmt"
	"sort"
)

func search(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target) // 返回搜索数的位置 target 第一个位置
	//fmt.Println(leftmost)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	//fmt.Println(rightmost)
	return rightmost - leftmost + 1
}

// 利用二分查找 找第一个位置，和最后一个位置
func searchV2(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	low := 0
	up := n - 1
	for low < up {
		mid := (low+up)/2 + 1
		if nums[mid] > target {
			up = mid - 1
		} else {
			low = mid
		}
	}
	right := low
	if nums[right] != target {
		return 0
	}
	fmt.Println(right)
	low = 0
	up = right
	for low < up {
		mid := (low + up) / 2
		fmt.Println(mid)
		if nums[mid] > target {
			up = mid - 1
		} else {
			low = mid
		}
	}
	left := low
	fmt.Println(right)
	return right - left + 1
}

func main() {
	arr := []int{1, 2, 2, 2, 4}
	fmt.Println(search(arr, 2))
	//fmt.Println(searchV2(arr,2))
}
