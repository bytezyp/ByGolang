package main

import "fmt"

func main() {
	m1 := make(map[int]int)
	fmt.Println(m1)
	m1[11] = 11
	fmt.Println(len(m1))
}
