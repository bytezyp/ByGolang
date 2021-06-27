package main

import "fmt"

func main() {
	m1 := make(map[int]int)
	fmt.Println(m1)
	m1[11] = 11
	m1[12] = 12
	m1[14] = 14
	m1[13] = 13

	fmt.Println(len(m1))
	for v, k := range m1 {
		fmt.Println(v, k)
	}
}
