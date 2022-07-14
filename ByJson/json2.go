package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type ConfigV3 struct {
	//Rules map[string][]map[string]interface{} `yaml:"rules"`
	AppConf *AppConfigV3 `yaml:"app"`
}

type AppConfigV3 struct {
	GdtConf      *AdxConfigV3 `yaml:"gdt"`
	OceanConf    *AdxConfigV3 `yaml:"ocean"`
	KuaishouConf *AdxConfigV3 `yaml:"kuaishou"`
	CpcConf      *AdxConfigV3 `yaml:"cpc"`
	IqiyiConf    *AdxConfigV3 `yaml:"iqiyi"`
	BaiduConf    *AdxConfigV3 `yaml:"baidu"`
	XiaomiConf   *AdxConfigV4 `yaml:"xiaomi"`
}
type AdxConfigV3 struct {
	QttBlacklistSwitch    bool                    `yaml:"qtt_blacklist_switch"`
	CaidanBlacklistSwitch bool                    `yaml:"caidan_blacklist_switch"`
	RecallCacheTime       int32                   `yaml:"recall_cache_time"`
	NewCacheTime          int32                   `yaml:"new_cache_time"`
	Bridge                BridgeConfig            `yaml:"bridge"`
	SwitchStrategies      []string                `yaml:"switch_strategies"`
	AppStrategies         []*App_Strategy_Item_V3 `yaml:"app_strageties"`
}

// xiaomi 优化
type AdxConfigV4 struct {
	QttBlacklistSwitch    bool                    `yaml:"qtt_blacklist_switch"`
	CaidanBlacklistSwitch bool                    `yaml:"caidan_blacklist_switch"`
	RecallCacheTime       int32                   `yaml:"recall_cache_time"`
	NewCacheTime          int32                   `yaml:"new_cache_time"`
	Bridge                BridgeConfig            `yaml:"bridge"`
	SwitchStrategies      []string                `yaml:"switch_strategies"`
	AppStrategies         []*App_Strategy_Item_V4 `yaml:"app_strageties"`
}
type App_Strategy_Item_V4 struct {
	StrategyUserIds  []string                   `yaml:"strategy_user_ids"`
	ShieldStrategies []*Shield_Strategy_Item_V3 `yaml:"shield_strageties"`
}

type App_Strategy_Item_V3 struct {
	StrategyUserIds  []int64                    `yaml:"strategy_user_ids"`
	ShieldStrategies []*Shield_Strategy_Item_V3 `yaml:"shield_strageties"`
}

type Shield_Strategy_Item_V3 struct {
	APP       string `yaml:"app"`
	Type      string `yaml:"type"`
	Option    string `yaml:"option"`
	Parameter string `yaml:"parameter"`
	Default   bool   `yaml:"default"` // 无设备投放控制 true 投放 false 不投放
}
type BridgeConfig struct {
	StrategyRate         uint64     `yaml:"strategy_rate"`
	RtaIds               []int64    `yaml:"rta_ids"`
	QttBridgeForBilibili BridgeInfo `yaml:"qtt_bridge_for_bilibili"`
	QttBridgeForKuaishou BridgeInfo `yaml:"qtt_bridge_for_kuaishou"`
}

type BridgeInfo struct {
	ReqRtaRate          uint64  `yaml:"req_rta_rate"`
	RtaId               int64   `yaml:"rta_id"`
	Url                 string  `yaml:"url"`
	TimeOut             int64   `yaml:"timeOut"`
	UserIds             []int64 `yaml:"user_ids"`
	AdvUserIds          []int64 `yaml:"adv_user_ids"`
	CacheExpTime        int64   `yaml:"cache_exp_time"`
	RefusedCacheExpTime int64   `yaml:"refused_cache_exp_time"`
	IsOverTimeBid       int32   `yaml:"is_over_time_bid"`
	T0ActiveFilter      bool    `yaml:"t0_active_filter"`
}

var GConfigV3 = new(ConfigV3)

func UnmarshalFile(name string, obj interface{}) error {
	content, err := ioutil.ReadFile(filepath.Join("/Users/zyp/go/ByGolang/ByJson/", name))
	if err != nil {
		return err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	return decoder.Decode(obj)
}
func main4() {
	UnmarshalFile("a.yaml", GConfigV3)
	fmt.Println(GConfigV3.AppConf.XiaomiConf.AppStrategies[0].StrategyUserIds[0])
	fmt.Printf("%T", GConfigV3.AppConf.XiaomiConf.AppStrategies[0].StrategyUserIds[0])
}
