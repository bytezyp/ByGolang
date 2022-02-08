package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 2
	change(slice...)
	fmt.Println(slice, len(slice), cap(slice))
	slice1 := slice[0:2]
	change(slice1...)

	fmt.Println(slice, len(slice), cap(slice))
}
