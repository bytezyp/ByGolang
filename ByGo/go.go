package main

import (
	"fmt"
	"time"
)

func demo()  {
	fmt.Println(333)
}

func main()  {
	//for i := 0; i < 5; i++ {
	//	//time.Sleep(time.Second)
	//	fmt.Println("Hello World")
	//}
	//runtime.GOMAXPROCS(1)
	//debug.SetMaxThreads(1)
	//var w sync.WaitGroup
	//w.Add(1)
	go func() {
		fmt.Println(123)
	}()
	//go func() {
	//	for  {
	//		demo()
	//	}
	//}()
	//w.Wait()
	time.Sleep(1*time.Second)
}
