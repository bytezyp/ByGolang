package main

import (
	"fmt"
	"math"
)

func sourBucket(nums []int,bucketNum int) []int  {
	// bucketNum
	numsLen := len(nums)
	// 确认数组范围
	min,max := nums[0],nums[0]
	for i := 0; i < numsLen; i++ {
		if min >nums[i] {
			min = nums[i]
		}
		if max < nums[i]{
			max = nums[i]
		}
	}
	bucket := make([][]int, bucketNum)
	for j := 0; j < numsLen; j++ {
		// 选择，分桶
		n := int(math.Floor(float64(nums[j] - min) / (float64(max - min + 1) / float64(bucketNum))))
		//fmt.Println(n,float64(nums[j] - min),float64(max - min + 1),float64(bucketNum),(float64(max - min + 1) / float64(bucketNum)))
		bucket[n] = append(bucket[n],nums[j])
	}
	fmt.Println(bucket)
	return nums
}

func main()  {
	arr := []int{1,3,77,44,4,33}
	fmt.Print(arr)
	sourBucket(arr,2)
}
