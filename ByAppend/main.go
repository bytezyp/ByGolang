package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	fmt.Printf("s1_pt=%p,s2_pt=%p \n", &s1, &s2)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
	fmt.Printf("s1_pt=%p,s2_pt=%p \n", &s1, &s2)
	a := make([]int, 0)
	fmt.Printf("a_pt=%p \n", &a)
	a = append(a, 1, 2, 3, 4, 5, 6)
	fmt.Printf("a_pt=%p \n", &a)
	var aa = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	// range 循环的副本，
	for i, v := range aa {
		if i == 0 {
			aa[1] = 12
			aa[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", aa)
}
