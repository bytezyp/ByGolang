package main

import "fmt"

func main1() {
	A := make(chan bool, 1)
	B := make(chan bool)
	Exit := make(chan bool)
	go func() {
		s := []int{1, 2, 3, 4}
		for i := 0; i < len(s); i++ {
			if ok := <-A; ok {
				fmt.Println("A: ", s[i])
				B <- true
			}
		}
	}()
	go func() {
		defer func() {
			close(Exit)
		}()
		s := []byte{'A', 'B', 'C', 'D'}
		for i := 0; i < len(s); i++ {
			if ok := <-B; ok {
				fmt.Printf("B: %c\n", s[i])
				A <- true
			}
		}
	}()
	A <- true
	<-Exit
}
