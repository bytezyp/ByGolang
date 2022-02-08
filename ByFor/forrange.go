package main

import "fmt"

func main() {
	arr := make(map[string][]int, 10)
	arr["aa"] = []int{1, 2}
	arr["bb"] = []int{3, 4}
	for _, value := range arr {
		//fmt.Println(value)
		for _, i2 := range value {
			if i2 == 3 {
				//continue
				break
			}
			fmt.Print(i2)
		}
	}
}
