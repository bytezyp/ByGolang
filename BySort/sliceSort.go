package main

import (
	"fmt"
	"math/rand"
	"sort"
)

var demo int

// 消息体
type FunnelData struct {
	AdId       uint64
	CreativeId uint64
	Ecpm       uint64
	Order      int
}
type EcpmFunnelData struct {
	Rand *rand.Rand
	Ad   []*FunnelData
}

func (p EcpmFunnelData) Len() int      { return len(p.Ad) }
func (p EcpmFunnelData) Swap(i, j int) { p.Ad[i], p.Ad[j] = p.Ad[j], p.Ad[i] }
func (p EcpmFunnelData) Less(i, j int) bool {
	if p.Ad[i].Ecpm > p.Ad[j].Ecpm {
		return true
	} else if p.Ad[i].Ecpm < p.Ad[j].Ecpm {
		return false
	} else {
		// 相同ecpm，随机选择
		return p.Rand.Intn(2) == 0
	}
}
func Print(fData []*FunnelData) {
	for _, datum := range fData {
		fmt.Printf("%v \n", datum)
	}
	println("-----------------")
}
func Init() {
	demo = 123
	fmt.Println(demo, "-----")
}
func main() {
	Init()
	fmt.Println(demo, "123")
	ss := make([]int, 0, 1)
	fmt.Println(ss)
	ss = append(ss, 1)
	fmt.Println(ss)
	return
	fData := make([]*FunnelData, 0, 4)
	for i := 0; i < 3; i++ {
		j := uint64(i)
		f := &FunnelData{AdId: j, CreativeId: j, Ecpm: j}
		fData = append(fData, f)
	}
	Print(fData)
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	sort.Sort(EcpmFunnelData{Rand: r3, Ad: fData})
	Print(fData)
	newDetail := make([]*FunnelData, 0, len(fData))
	for i, dt := range fData {
		fData := &FunnelData{
			AdId:       dt.AdId,
			CreativeId: dt.CreativeId,
			Ecpm:       dt.Ecpm,
			Order:      i,
		}
		newDetail = append(newDetail, fData)
	}
	Print(newDetail)
}
