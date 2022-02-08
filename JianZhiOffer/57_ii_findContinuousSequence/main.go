package main

import "fmt"

func findContinuousSequence(n int) [][]int {
	mid := n/2 + 1
	ret := [][]int{}
	for i := 1; i <= mid; i++ {
		sum := 0
		arr := []int{}
		for j := i; j <= mid; j++ {
			sum = sum + j
			arr = append(arr, j)
			if sum < n && j == mid {
				arr = []int{}
				break
			} else if sum < n {
				continue
			} else if sum > n {
				arr = []int{}
				break
			} else {
				break
			}

		}
		if len(arr) > 0 {
			ret = append(ret, arr)
		}
	}
	return ret
}

func main() {
	fmt.Println(findContinuousSequence(9))
}
