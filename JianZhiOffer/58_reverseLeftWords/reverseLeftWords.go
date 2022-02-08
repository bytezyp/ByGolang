package main

import "fmt"

func reverseLeftWord(str string, n int) string {
	bs := []byte(str)
	var reverse func(start, end int)
	reverse = func(start, end int) {
		for start < end {
			bs[start], bs[end] = bs[end], bs[start]
			start++
			end--
		}
	}
	fmt.Println(string(bs))
	reverse(0, len(bs)-1)
	fmt.Println(string(bs))
	reverse(0, len(bs)-n-1)
	fmt.Println(string(bs))
	reverse(len(bs)-n, len(bs)-1)
	fmt.Println(string(bs))
	return string(bs)
}

func reverseLeftWords(s string, n int) string {
	if len(s) <= 0 {
		return s
	}
	bs := []byte(s)
	bs1 := bs[n:]
	bs2 := bs[:n]
	newBs := append(bs1, bs2...)
	return string(newBs)
}

func main() {
	//fmt.Println(reverseLeftWord("12345", 2))

	fmt.Println(reverseLeftWords("12345", 3))
}
