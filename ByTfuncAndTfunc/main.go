package main

import (
	"ByGolang/ByTfuncAndTfunc/gom"
	"fmt"
)

func main() {
	point := gom.Point{}
	point.SetX(1)
	fmt.Println(point.X())
	point2 := &gom.Point{}
	point2.SetX(888)
	fmt.Println(point2.X())
	person := &gom.Person{"zhang"}
	fmt.Println(person.GetName())
}
