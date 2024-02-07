package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"time"
)

func main() {
	println(TodayTimestamp())
	t := time.Now().AddDate(0, 0, -7).Format("2006-01-02 00:00:00")
	fmt.Println(t)
	//IsHitExperiment("123", 1)
}

func TodayTimestamp() int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	dt, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), loc)
	return dt.Unix()
}

func IsHitExperiment(Uid string, factor uint64) (exp int) {
	exp = 1
	didHash := murmur3.Sum64([]byte(Uid))
	if factor == 0 {
		factor = 10
	}
	if didHash%1000 < factor {
		return exp
	}
	return exp
}

func GenCode(weight int) {
	//rand.In
}
