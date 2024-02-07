package main

import (
	"fmt"
	"gitlab.ushareit.me/ads/architecture/golib.git/glog"
)

func main() {
	var a uint64 = 123

	glog.Infof("echo : %d", a)
	m := make(map[int]int64, 2)
	fmt.Println(len(m))
	for i, i2 := range m {
		fmt.Println(i, i2)
	}
}
