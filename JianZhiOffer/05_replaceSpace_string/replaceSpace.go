package main

import "fmt"

func replaceSpace(str string) string {
	var replace []byte
	//
	for _, v := range []byte(str) {
		if v == ' ' {
			replace = append(replace, []byte(`%20`)...)
			continue
		}
		replace = append(replace, v)
	}
	return string(replace)
}

func main() {
	fmt.Println(replaceSpace("i am   zyp"))
}
