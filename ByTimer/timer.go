package main

import (
	"fmt"
	"strconv"
)

func main() {
	//timer := time.NewTimer(2 * time.Second)
	//defer timer.Stop()
	//go func(timer *time.Timer) {
	//	<-timer.C
	//	println(222)
	//}(timer)
	//fmt.Println("测试两秒时间")

	res := ""
	//println(res)
	vs := []string{"1", "2", "3"}
	var lowerBound uint64
	lowerBound = 1000
	for _, v := range vs {
		uv, e := strconv.ParseUint(v, 10, 64)
		if e != nil {

		}
		res = fmt.Sprintf("%s,%d", res, uv+lowerBound)
	}
	println(res)
	println(res[1:])
}
