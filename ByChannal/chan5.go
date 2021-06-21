package main

import (
	"fmt"
	"time"
)

func main5() {
	exit := make(chan int)
	defer func() {
		exit <- 1000
	}()
	ch := getChan()
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
			time.Sleep(2 * time.Second)
		}
	}
	fmt.Println(<-exit)
}

func getChan() chan int {
	ch := make(chan int, 10)
	i := 1
	go func() {
		for {
			ch <- i
			fmt.Println("chan len:", len(ch))
			i++
		}
	}()
	return ch
}
