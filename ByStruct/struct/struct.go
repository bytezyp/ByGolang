package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) GetSelf() *Person {
	return p
}

func (p Person) Clone() *Person {
	return &Person{}
}

func PrintPerson() {
	pp := &Person{}
	fmt.Printf("%p \n", pp.GetSelf())
	fmt.Printf("%p", pp.Clone())
}
func main() {
	PrintPerson()
	return
	m := make(map[int]string)
	m[1] = "11"
	m[2] = "22"

	for i := range m {
		fmt.Println(i)
	}
	slice := make([]int, 3)
	for k := range slice {
		fmt.Println(k)
	}
}
