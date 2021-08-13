package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//m1 := make(map[int]int)
	//fmt.Println(m1)
	//m1[11] = 11
	//m1[12] = 12
	//m1[14] = 14
	//m1[13] = 13
	//
	//fmt.Println(len(m1))
	//for v, k := range m1 {
	//	fmt.Println(v, k)
	//}

	m2 := make(map[string][]int, 1470000)
	arr := make([]int,0,10)
	fmt.Printf("slice size: %d\n", unsafe.Sizeof(arr))
	for i := 0; i < 100000; i++ {
		arr = append(arr, 10296005)
	}
	//fmt.Println(arr)
	fmt.Printf("slice size: %d\n", unsafe.Sizeof(arr))
	fmt.Println("Size of [200]int:", unsafe.Sizeof([200]int{}))
	fmt.Println("Size of [180]int:", unsafe.Sizeof([180]int{}))
	m2["b95fb94c892f4051ab3905033ff48574"] = arr
	m2["b95fb94c892f4051ab3905033ff48571"] = arr
	//fmt.Println(m2)
	//fmt.Printf("map size: %d\n", unsafe.Sizeof(m2["b95fb94c892f4051ab3905033ff48571"]))
	//fmt.Printf("map size: %d\n", unsafe.Sizeof(m2["b95fb94c892f4051ab3905033ff48574"]))
}
