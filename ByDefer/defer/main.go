package main

import "fmt"

var g = 100

func f() (r int) {
	r = g
	defer func() {
		r = 200
	}()
	//r = 0
	return r
}
func t() int {
	defer func() {
		g = 200
	}()
	return g
}
func main() {
	//i := f()
	//fmt.Printf("main: i= %d , g = %d \n", i, g) // i = 200  g = 100
	i := t()
	fmt.Printf("main: i= %d , g = %d \n", i, g) // i = 100  g = 200
}
