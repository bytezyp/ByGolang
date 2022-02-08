package main

import (
	"fmt"
	"time"
)

func TodayTimestamp() int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	dt, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), loc)
	return dt.Unix()
}
func main() {
	print(TodayTimestamp())
	var layer int64
	fmt.Println("int64:", layer)
	return
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
	var cc bool
	aa := false
	bb := true
	cc = aa || bb
	fmt.Println(cc)
	cc = aa && bb
	fmt.Println(cc)
}
