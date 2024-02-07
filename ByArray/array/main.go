package main

import (
	"fmt"
)

func arrayToSlice()  {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3]
	//rand.Shuffle()
	fmt.Println(slice)
}

func main()  {
	arrayToSlice()
}
