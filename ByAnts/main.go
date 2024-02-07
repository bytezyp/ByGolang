package main

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	json "github.com/json-iterator/go"
	"time"
)

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

func ret(n int) {
	now := time.Now().AddDate(0, 1, 2).Format("2006-01-02")
	fmt.Println(now)
}
func TestMarshal() {
	m := make(map[string]string)
	m["a"] = "sss&aaa"

	s, _ := json.Marshal(m)
	fmt.Println(string(s))
}

func docreative() {
	arr := []string{
		"pangle",
		"meitu",
		"mytarget",
		"smaato",
		"pubnative",
		"criteo",
		"pubnative_low",
		"smaato_low",
		"smaato_high",
		"smadex",
		"yandex_shareit",
		"telecoming",
		"bigo",
		"RTBhouse_APEC",
		"RTBhouse_EMEA",
		"webeye_pangle_SHAREit",
		"webeye_pangle_US_SHAREit",
		"pubnative_GBwhatsapp",
		"smaato_GBwhatsapp",
		"yeahmobi_GBwhatsapp",
		"smadex_meetchat",
		"smadex_vml",
		"telecoming_vml",
		"yeahmobi_vml",
		"pubnative_vml",
		"smaato_vml",
		"smaato_slite",
		"pubnative_slite",
		"smadex_slite",
		"bigo_slite",
		"yeahmobi_slite",
		"pangle_slite",
		"yeahmobi_Zshare",
		"bigo_Zshare",
		"pubnative_Zshare",
		"smaato_Zshare",
		"smaato_zslite",
		"yeahmobi_zslite",
		"bigo_zslite",
		"pubnative_zslite",
		"yeahmobi_online",
		"opera_shareit_ru",
		"xapads_sg_shareit",
		"adsnowy_shareit_eu_native",
		"adsnowy_shareit_eu_video",
		"smaato_nopos",
		"smadex_nopos",
		"pubmatic",
		"yeahmobi_nopos",
		"pubnative_nopos",
		"telecoming_nopos",
		"opera_nopos",
		"yandex_nopos",
		"loopme_shareit",
	}
	newStr := "insert into geoedge_ad_info(project_id,project_name,dsp,country,cid,crid,scan_id,bid_cat,risk_status,status) values "
	for _, str := range arr {
		uuid, _ := uuid2.NewUUID()
		key := uuid.String()
		val := fmt.Sprintf("('%s','%s_DZ_45053191','%s','DZ','45053191','','','',2,2)", key, str, str)
		newStr += "," + val
	}
	fmt.Println(newStr)
}

func main() {
	docreative()

	return
	currentTime := time.Now().Format("2006-01-02")
	fmt.Println(currentTime)
	var num int64
	num = 12
	if num == 12 {
		fmt.Println(11111)
	}
	//var str []string{'10.71.0.181:6379,10.71.0.182:6379'}
	//var BiddingLogConfigMap sync.Map
	//v, _ := BiddingLogConfigMap.Load("key")
	//fmt.Println(v)

	timeStr := time.Now().Format("2006-01-01")
	fmt.Println(timeStr)
	//mp := make(map[int]int, 10)
	//mp[33] = 33
	//a := 33
	//var v int
	//if v, ok := mp[a]; !ok {
	//	println(v)
	//}
	//println(v)
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
