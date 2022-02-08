package main

import (
	"fmt"
	"math"
)

func strToInt(str string) int {
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	str = str[i:]

	ans := 0      //64位数来存储好比较
	flag := false //符号位

	for i, v := range str {
		if v >= '0' && v <= '9' {
			// 这里防止前边 有 0
			// 字符转数字： “此数字的 ASCII 码” 与 “ 00 的 ASCII 码” 相减即可；
			// 所以全部减去 '0'的ASCII
			fmt.Println(int(v-'0'), v, v-'0')
			ans = ans*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			flag = true
		} else if v == '+' && i == 0 {
		} else {
			break
		}

		if ans > math.MaxInt32 {
			if flag {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	if flag {
		return -1 * ans
	}
	return ans

}
func main() {
	fmt.Println(strToInt(" -099"))
	str := "abc"
	// 遍历字符串，是遍历 打印 他的阿斯科马值
	// 字符转数字： “此数字的 ASCII 码” 与 “ 00 的 ASCII 码” 相减即可；
	for i, v := range str {
		fmt.Println(v, i)
	}
}
