package main

import (
	"fmt"
	"time"
)

type AA struct {
	s string
}

// 这是上面提到的 "在方法内把局部变量指针返回" 的情况
func foo(s string) *AA {
	a := new(AA)
	a.s = s
	return a //返回局部变量a,在C语言中妥妥野指针，但在go则ok，但a会逃逸到堆
}
func main() {
	//a := foo("hello")
	//b := a.s + " world"
	//c := b + "!"
	//fmt.Println(c)

	work := make(chan int, 10)
	for i := 0; i < 2; i++ {
		go func(cc chan int) {
			for true {
				select {
				case num := <-cc:
					fmt.Println(num)
				default:
					//fmt.Println("default")

				}
			}

		}(work)
	}

	go func() {
		for i := 1; i <= 1000; i++ {
			work <- i
		}
	}()
	time.Sleep(10 * time.Second)
}
