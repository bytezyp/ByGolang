package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)          // 获取原始数据长度
	sort.Ints(nums)         // 数组排序
	ans := make([][]int, 0) // 定义返回值
	fmt.Println(ans)
	// 从排序的数组中，最小的开始
	for first := 0; first < n; first++ {
		// 需要和上一次的枚举不一样
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1             // 定义第三个指针开始
		target := -1 * nums[first] // a+b = target
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 左右指针相等 及跳出本次循环
			if second == third {
				break
			}
			if nums[third]+nums[second] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))
}
