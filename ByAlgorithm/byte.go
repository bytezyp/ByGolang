package main

import (
	"ByGolang/ByAlgorithm/slicestacks"
	"fmt"
	"unsafe"
)

type St1 struct {
	A byte
	B int
	C byte
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			fmt.Println(s[rk+1])
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	fmt.Println(m)
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	//nn := lengthOfLongestSubstring("abcdder")
	//fmt.Println(nn)
	m := map[byte]int{}
	str := "abc"
	m[str[0]]++
	fmt.Printf("%v", m)
	fmt.Printf("%v", m[10])
	fmt.Printf("%v %T", str[0], str[0])
	return
	stack := slicestacks.New(11, 22)
	fmt.Println(stack.Size())
	fmt.Printf("%+v", stack)
	fmt.Println("st1 占用的字节数：" + fmt.Sprint(unsafe.Sizeof(St1{}.A)))
	fmt.Println("st1 对齐的字节数：" + fmt.Sprint(unsafe.Alignof(St1{}.B)))
	fmt.Println("ST1结构体 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(St1{})))
	fmt.Println("ST1结构体 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(St1{})))
}
