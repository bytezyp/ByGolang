package main

import (
	"fmt"
	"time"
)

func main6() {
	ch := make(chan int, 10)

	// 读取chan
	go func() {
		for {
			select {
			case i := <-ch: // 读取到15个的时候，就一直堵塞的，然后就一直走default
				// 只读取15次chan
				fmt.Println(i)
			default:
				// 读取15次chan以后的操作一直在这个空语句无任何IO操作的default条件里死循环，无法出让P，以保证一个GPM关系。
				// 而如果无default条件的话，则系统当读取完15次chan后，当前goroutine会发生 chan IO 阻塞, Go调度器根据GPM的调度关系,会将当前执行关系中的G切换出去，再从LRQ队列中取一个新的G，重新组成一个GPM继续执行，以实现合理利用计算机资源，提高GO的高并发性能
			}
		}
	}()

	// 写入10个值到chan
	for i := 0; i < 15; i++ {
		ch <- i
	}

	// 模拟程序效果使用
	time.Sleep(time.Minute)
}
