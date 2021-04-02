package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=1; i<5;i++ {
		time.Sleep(time.Second)
		fmt.Println("hello gmp")
	}
}
