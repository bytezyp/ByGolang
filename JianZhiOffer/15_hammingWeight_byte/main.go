package main

import "fmt"

func hammingWeight(num uint32) (n int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			n++
		}
	}
	return
}

func hammingWeight2(num int) (res int) {
	for num > 0 {
		//fmt.Println(num&1)
		res += num & 1
		num >>= 1
	}
	return res
}

func main() {
	//fmt.Println(hammingWeight(5))
	fmt.Println(hammingWeight2(5))
}
