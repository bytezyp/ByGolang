package main

import "fmt"

type FunnelData struct {
	AdId   int
	AdType string
	AdBid  float64
	Ecpm   string
}

type FunnelInfo struct {
	System      string // 系统名称
	Name        string // 漏斗名称
	InputAdNum  int    // 进入的数量
	OuputAdNum  int    // 输出的数量
	FilterAdNum int    // 过滤的条数
	DetailType  int    // 记录详情类型
	Detail      []*FunnelData
}

func main() {
	mp := make(map[int]int, 10)
	mp[33] = 33
	a := 33
	var v int
	if v, ok := mp[a]; !ok {
		println(v)
	}
	println(v)
	return
	//n := runtime.GOMAXPROCS(0) //指定以2核运算
	//fmt.Println(runtime.NumCPU())
	//fmt.Println("n = ", n)
	//
	//for {
	//	go fmt.Print(1)
	//
	//	fmt.Print(0)
	//	time.Sleep(1 * time.Second)
	//}
	m := make([]int, 0)
	fmt.Println(cap(m))
	m = append(m, 1)
	fmt.Println(cap(m))
	m = append(m, 1)
	fmt.Println(cap(m))
	var userId string
	need := userId != ""
	fmt.Println(11111, need)
}
