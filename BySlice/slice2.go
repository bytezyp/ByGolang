package main

import "fmt"

func main2() {
	s := make([]int, 0, 1)
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("s: len:%d , cap:%d \n", len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("append s: len:%d , cap:%d \n", len(s), cap(s))
	s1 := append(s, 4, 5, 6)
	fmt.Printf("append s1: len:%d , cap:%d", len(s1), cap(s1))
}
