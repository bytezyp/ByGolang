package main

import "fmt"

func modifySlice(innerSlice []string) {
	fmt.Printf("%p %v   %p\n", &innerSlice, innerSlice, &innerSlice[0])
	innerSlice = append(innerSlice, "a")
	fmt.Println(len(innerSlice), cap(innerSlice))
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
	fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, &innerSlice[0])
}

func modifySliceV2(innerSlice *[]string) {

	fmt.Printf("%p %v   %p\n", &innerSlice, innerSlice, &innerSlice)
	*innerSlice = append((*innerSlice), "a")
	fmt.Println(len(*innerSlice), cap(*innerSlice))
	//innerSlice[0] = "b"
	//innerSlice[1] = "b"
	//fmt.Println(innerSlice)
	//fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, &innerSlice[0])
}

//  go tool compile -m -l -m  main.go
func main() {
	//a := 666
	//fmt.Println(a)
	outerSlice := []string{"a", "a"}

	//outerSlice := make([]string,0,2)
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice)
	outerSlice = append(outerSlice, "1", "2", "3")
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice)
	outerSlice = append(outerSlice, "3")
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice)
	return
	fmt.Println(len(outerSlice), cap(outerSlice))
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice[0])
	modifySlice(outerSlice)
	fmt.Println(outerSlice)
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice[0])
	modifySliceV2(&outerSlice)
	fmt.Println(len(outerSlice), cap(outerSlice))
}
