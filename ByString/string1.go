package main

import "fmt"

const a = 1 << iota
const b = iota

func main() {
	fmt.Println(a, b)
}
