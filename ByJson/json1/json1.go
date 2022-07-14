package main

import (
	"encoding/json"
	"fmt"
)

// 消息体
type FunnelData struct {
	AdId uint64
	Ecpm uint64
}

// 子阶段
type FunnelChildInfo struct {
	Name   string
	AdNum  int
	Detail []*FunnelInfo
}

// 广告位相关
type FunnelPostInfo struct {
	PostId      uint64 // 广告位id
	CountryCode string // 国家
	AppId       string // app名称
}

// 系统阶段漏斗
type FunnelInfo struct {
	System      string // 系统名称
	Name        string // 漏斗名称
	InputAdNum  int    // 进入数量
	OutputAdNum int    // 输出数量
	FilterAdNum int    // 过滤了多少
	PostInfo    *FunnelPostInfo
	Child       map[string]*FunnelChildInfo
	Detail      []*FunnelData
}

func main() {
	funnel := make(map[string]*FunnelInfo, 1)
	post := &FunnelPostInfo{
		PostId:      1,
		CountryCode: "zh",
		AppId:       "shareit",
	}
	detail := make([]*FunnelData, 0, 1)
	detail = append(detail, &FunnelData{
		AdId: 123,
		Ecpm: 0,
	})
	funnel["index"] = &FunnelInfo{
		System:      "retrieval",
		Name:        "index",
		InputAdNum:  22,
		OutputAdNum: 22,
		FilterAdNum: 0,
		PostInfo:    post,
		Child:       nil,
		Detail:      detail,
	}
	jsonData, _ := json.Marshal(funnel["index"])
	fmt.Println(string(jsonData))
}
