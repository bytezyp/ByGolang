package main

import (
	"fmt"
)

type Cat struct {
	age     int
	name    string
	friends []int
}

func main() {
	wilson := Cat{7, "Wilson", []int{11, 22, 33}}
	nikita := wilson
	//fmt.Println(nikita)
	nikita.friends = make([]int, len(wilson.friends))
	fmt.Println(nikita)
	copy(nikita.friends, wilson.friends)
	wilson.friends[0] = 66
	//nikita.friends = append(nikita.friends, 44)
	fmt.Println(wilson)
	fmt.Println(nikita)
}
