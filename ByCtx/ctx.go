package main

import "fmt"

type Data struct {
	Name   string
	Id     int
	Person map[int]string
}
type Ctx struct {
	D *Data
}

func main() {
	c := &Ctx{
		D: &Data{},
	}
	c.D.Name = "zhang"
	c.D.Id = 1
	c.D.Person = make(map[int]string, 5)
	c.D.Person[1] = "111"
	fmt.Println(c.D.Name, c.D.Id, c.D.Person)

	person := c.D.Person
	person[1] = "222"
	fmt.Println(c.D.Name, c.D.Id, person, c.D.Person)
	var PixalateRate float64
	if PixalateRate <= 0 {
		println(1111)
	}
	//println(PixalateRate)

	if true && true || false {
		println(333)
	}
}
