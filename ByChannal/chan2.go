package main

import "fmt"

func main2() {
	ch := make(chan int)
	exit := make(chan struct{})

	go func() {
		for i := 1; i <= 20; i++ {
			//println("go1:start")

			println("g1:", <-ch) // 执行步骤1， 执行步骤5
			i++                  //执行步骤6
			//fmt.Println(123)
			ch <- i // 执行步骤7
		}
	}()

	go func() {
		defer func() {
			close(ch)
			close(exit)
		}()
		for i := 0; i < 20; i++ {
			//println("go2:start")
			i++                  // 执行步骤2
			ch <- i              //执行步骤3
			println("g2:", <-ch) //执行步骤4
		}
	}()
	//time.Sleep(time.Second*1)
	//<-exit
	fmt.Println(<-exit)
}
