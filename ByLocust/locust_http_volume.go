package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/myzhan/boomer"
)

func main() {
	task := &boomer.Task{
		Name: "sostreq",
		// The weight is used to distribute goroutines over multiple tasks.
		Weight: 10,
		Fn:     postDemo,
	}
	boomer.Run(task)

}

func postDemo() {
	start := time.Now()
	dataBody := `{
    "reserved_app": [
        {
            "package_name": "i1d.co.reliancerobopds.mobile",
            "md5sum": "78b7614003cca6297ddea421a28368b4",
            "cid": 72002
        },
        {
            "package_name": "1org.mozilla.firefox_beta",
            "md5sum": "78b7614003cca6297ddea421a28368b4",
            "cid": 73276
        }
    ],
    "new_cache_request": 1,
    "placements": [
        {
            "pos_id": 1643,
            "ad_count": 15,
            "ad_bottom_count": 1,
            "ad_offline_count": 1,
            "ad_cache_count": 1,
            "ad_startup_count": 1,
            "support_video": true,
            "load_type": 0,
            "is_keep_popup": 0,
            "lp_package": ""
        }
    ],
    "existed_ad": [],
    "app_info": {
        "app_pkg": "news.buzzfeed.buzznews",
        "app_ver": 4150633,
        "app_vername": "6.9.1",
        "channel": "GOOGLEPLAY",
        "app_key": "09d24136-d3a0-47b9-bf0b-b4662a01ee53",
        "i_ms": 1254360,
        "u_ms": 1254360,
        "init_time": 1623503297996,
        "sdk_ver": 4080633
    },
    "target": {
        "forced_country": "",
        "forced_city": "",
        "pkgs": ""
    },
    "user": {
        "device_id": "m.A2B2EE711232",
        "user_id": "6joJTR",
        "beyla_id": "b8fc0ca638c84e1d89c8deee2663fe2w",
        "limit_ad_tracking": false
    },
    "device_info": {
        "geo": {
            "lat": 33,
            "lon": 36,
            "station": ""
        },
        "device_type": "phone",
        "os_type": "android",
        "os_ver": 29,
        "screen_width": 1080,
        "screen_height": 2183,
        "brand": "samsung",
        "manufacturer": "samsung",
        "device_model": "SM-A715F",
        "dpi": 420,
        "network": "WIFI",
        "mac": "A2B1EE711035",
        "imsi": "",
        "gaid": "52b8367b-9caa-3843-91df-17689c854e12",
        "cpu_bit": "64",
        "android_id": "9a2caebc753a5256",
        "timezone": "GMT+08:00",
        "lang": "zh",
        "country": "ID",
        "ip": "114.79.13.239",
        "battery_info": "BatteryInfo{batteryPercent=86, isUsbCharge=true, isAcCharge=false}",
        "cpu_abi": [
            "arm64-v8a",
            "armeabi-v7a",
            "armeabi"
        ]
    },
    "ext": {
        "gdpr_consent": true,
        "support_mraidjs": 1
    },
    "layer_config_version": "46f9c00efe1379c92ef2774e04d75e95",
    "ts": 1623504539999,
    "enable_action_tracker": 1,
    "rid": "",
    "debug": 1
}`
	resp, _ := http.Post("http://midas-api-volume.hellay.net/shareit/get_ads",
		"text/plain",
		strings.NewReader(dataBody))

	//defer resp.Body.Close()
	_, _ = ioutil.ReadAll(resp.Body)
	//fmt.Println(string(resp.Status), string(resp.StatusCode))
	elapsed := time.Since(start)
	if resp.Status == "200 OK" {
		boomer.RecordSuccess("http", "sostreq", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
	} else {
		boomer.RecordFailure("http", "sostreq", elapsed.Nanoseconds()/int64(time.Millisecond), "sostreq not equal")
	}

	//fmt.Println(string(body))
}
