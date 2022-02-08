package main

import (
	"fmt"
	"math"
)

func reverse(num int) (rev int) {
	for num != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := num % 10 // 取余数
		num /= 10         // 取商
		fmt.Println(num, digit)
		rev = rev*10 + digit
	}
	return
}

func main() {
	//fmt.Println(reverse(-123))
	a := 0
	b := 123
	c := a | b
	fmt.Println(c)
}
