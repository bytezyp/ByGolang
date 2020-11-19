package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	slice := arr1[:2]
	fmt.Println(slice)
	for i := 1; i <= 8; i++ {
		slice = append(slice, 1)
	}
	fmt.Println(slice, arr1)
}
