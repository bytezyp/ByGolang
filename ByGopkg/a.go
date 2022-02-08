package main

import (
	"fmt"
	"github.com/bytedance/gopkg/lang/mcache"
)

func main() {
	ss := mcache.Malloc(1, 1)
	fmt.Printf("%v", ss)
	fmt.Printf("%v", 11)
}
