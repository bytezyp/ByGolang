package main

import "sync"

type Name struct {
	Val  int
	Lock sync.Mutex
}

func main() {

	go func() {

	}()
}
