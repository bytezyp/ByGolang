package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	max := 0
	j := -1
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(m, s[i-1])
		}
		// 正确的
		for j+1 < n && m[s[j+1]] == 0 {
			//println(j,string(s[j+1]))
			m[s[j+1]]++
			j++
		}

		max = getMax(max, j-i+1)
	}
	return max
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	str := "dvdf"
	fmt.Println(lengthOfLongestSubstring(str))
}
