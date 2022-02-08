package main

import (
	"fmt"
	"strconv"
)

func compressString(str string) string {
	if len(str) == 0 {
		return ""
	}
	length := len(str) - 1
	newStr := ""
	count := 0
	for i := 0; i <= length; i++ {
		s := str[i]
		if i > 0 {
			c := str[i-1]
			if s != c {
				newStr += string(c) + strconv.Itoa(count)
				count = 1
			} else {
				count++
			}

		} else {
			count++
		}
		// 做最后边界处理
		if i == length {
			newStr += string(s) + strconv.Itoa(count)
		}
	}
	return newStr
}
func main() {
	fmt.Println(compressString("aacbbc"))
}
