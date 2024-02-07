package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func dispersed(numsArr []int64, max int64) int {
	index := 0
	randRes, _ := rand.Int(rand.Reader, big.NewInt(max))
	randNum := randRes.Int64()
	arr := make([]int, 0, len(numsArr))
	for k, num := range numsArr {
		// 如果num为0，说明概率为0
		if num == 0 {
			continue
		}
		// 差值为0，说明正好命中
		diff := num - randNum
		if diff == 0 {
			index = k
			return index
		}
		// 大于0，说明符合取值的概率
		if diff > 0 {
			arr = append(arr, k)
		}
	}
	length := len(arr)
	if length == 0 {
		index = arr[0]
	}
	if length > 0 {
		indexRes, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		indexRand := indexRes.Int64()
		index = arr[int(indexRand)]
	}
	return index
}

func main() {
	//nums := []int64{1, 2, 38, 55, 19}

	for i := 0; i < 100; i++ {
		//fmt.Println(dispersed(nums, 55))
	}
	var err error
	fmt.Println(err)
	err = fmt.Errorf("%s", "123")
	fmt.Println(err)
}
