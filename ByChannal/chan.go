package main

import "fmt"

func main0() {
	ch := make(chan int)
	go func() {
		fmt.Println("go start")
		ch <- 1
		fmt.Println("go end")
	}()
	fmt.Println("all start ")
	<-ch
	fmt.Println("all end")
}
