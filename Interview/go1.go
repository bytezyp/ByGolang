package main

import "fmt"

func main1() {
	defer_call()
}
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触犯异常")
}

// defer 是后进先出，当遇到panic 时，会按照defer后进先出的执行，最后执行panic
