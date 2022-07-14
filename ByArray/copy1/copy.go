package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
}

func main() {
	// 浅拷贝和深拷贝
	// 简单来说，浅拷贝，是指两个变量，指向是同一个堆对象。
	// 深拷贝，是指两个变量，指向堆上的两个不同的堆对象。
	persons := make([]*Person, 0, 5)
	for i := 1; i <= 5; i++ {
		persons = append(persons, &Person{Name: "aa" + strconv.Itoa(i)})
	}
	for i, person := range persons {
		a := i + 1
		fmt.Printf("persons : %v %v \n", a, person)
	}
	fmt.Println("---------------------------- \n")
	persons1 := make([]*Person, 5)
	//fmt.Printf("%v %v", persons, persons1)
	copy(persons1, persons)
	//fmt.Printf("%v %v", persons, persons1)
	persons[3].Name = "7777"
	for i, person := range persons1 {
		a := i + 1
		fmt.Printf("persons : %v %v \n", a, person)
	}
}
