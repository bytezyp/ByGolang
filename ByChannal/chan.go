//package main
//
//import "fmt"
//
//func main() {
//	ch := make(chan int)
//	go func() {
//		fmt.Println("go start")
//		ch <- 1
//		fmt.Println("go end")
//	}()
//	fmt.Println("all start ")
//	<-ch
//	fmt.Println("all end")
//}

package main

import "fmt"

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println(123)
				return
			default: // ...
				fmt.Println("default start")
			}
		}
	}()
	// ...
	fmt.Println("all start")
	quit <- true
}
