package main

import "fmt"

func main() {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
	var  cc bool
	aa := false
	bb := true
	cc = aa || bb
	fmt.Println(cc)
	cc = aa && bb
	fmt.Println(cc)
}
