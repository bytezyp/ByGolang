package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}
func demo(i *int) {
	fmt.Println(*i)
}
func main() {
	i := 5
	defer hello(i)
	// 5
	i = i + 10
	defer hello(i)
	// 15
	// 指针和值拷贝，结果不一样，注意
	defer demo(&i)
}
