package main

import "fmt"

func getIndex(nums []int,num int)int  {
	var start int;
	end := len(nums)-1
	for start <= end {
		mid := start+(end-start)/2
		//fmt.Println(mid,start,end,mid)
		if nums[mid] == num {
			return mid
		}else if num < nums[mid] {
			end = mid-1
		}else {
			start = mid+1
		}
	}
	return -1
}

func main()  {
	arr := []int{1,3,4,33,44,777}
	fmt.Print(arr)
	fmt.Println(getIndex(arr,4))
}



