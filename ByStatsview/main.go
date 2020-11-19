package main

import (
	"github.com/go-echarts/statsview"
	"math/rand"
	"strconv"
	"time"
)

func work() {
	m := map[string][]byte{}
	for {
		b := make([]byte, 512+rand.Intn(16*1024))
		m[strconv.Itoa(len(b)%(10*100))] = b
		if len(b)%(10*100) == 0 {
			m = make(map[string][]byte)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
func main() {
	go work()
	mgr := statsview.New()
	mgr.Start()
}
