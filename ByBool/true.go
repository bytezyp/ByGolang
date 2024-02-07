package main

import (
	"encoding/base64"
	"fmt"
	"gitlab.ushareit.me/ads/architecture/golib.git/encodex"
	"net/url"
	"strconv"
	"time"
)

type TrackingData struct {
	ProcessTime     int64  `json:"process_time"`
	BidTime         int64  `json:"bid_time"`
	BidIp           string `json:"bid_ip"`
	ImpIp           string `json:"imp_ip"`
	RequestUri      string `json:"request_uri"`
	Referer         string `json:"referer"`
	Ua              string `json:"ua"`
	Platform        string `json:"platform"`
	PlatformId      string `json:"platform_id"`
	Country         string `json:"country"`
	Os              string `json:"os"`
	Osv             string `json:"osv"`
	BidId           string `json:"bid_id"`
	ImpId           string `json:"imp_id"`
	Udid            string `json:"udid"`
	Pid             string `json:"pid"`
	AdType          int    `json:"ad_type"`
	DeviceType      string `json:"device_type"`
	NetworkType     string `json:"network_type"`
	TrafficType     string `json:"traffic_type"`
	Width           string `json:"width"`
	Height          string `json:"height"`
	CSize           string `json:"c_size"`
	Gaid            string `json:"gaid"`
	DspId           string `json:"dsp_id"`
	DspName         string `json:"dsp_name"`
	Adomain         string `json:"adomain"`
	Crid            string `json:"crid"`
	Bundle          string `json:"bundle"`
	ShBidfloor      int64  `json:"sh_bidfloor"`
	Bidfloor        int64  `json:"bidfloor"`
	ShAdmPrice      int64  `json:"sh_adm_price"`
	Price           int64  `json:"price"`
	ShPrice         int64  `json:"sh_price"`
	AuctionPrice    int64  `json:"auction_price"`
	AuctionPriceRaw string `json:"auction_price_raw"`
	ClickUrl        string `json:"click_url"`
	Nurl            string `json:"nurl"`
}

func genData() map[string]string {
	data := make(map[string]string)
	data["bid_time"] = "1662534125"
	data["bid_ip"] = "127.0.0.1"
	data["imp_ip"] = "127.0.0.1"
	data["request_uri"] = ""
	data["referer"] = ""
	data["ua"] = ""
	data["platform"] = "vangle"
	data["country"] = "AM"
	data["os"] = "iOS"
	data["osv"] = "29"
	data["bid_id"] = "6316f966b41b7a94ac99b93c2"
	data["imp_id"] = "6316f966b41b7a94ac99b93c2"
	data["udid"] = "9a2caebc753a5256"
	data["pid"] = "123"
	data["ad_type"] = "2"
	data["device_type"] = "1"
	data["network_type"] = "4"
	data["traffic_type"] = "2"
	data["width"] = "400"
	data["height"] = "600"
	data["c_size"] = "1200*628"
	data["gaid"] = "52b8367b-9caa-3843-91df-17689c854e12"
	data["platform_id"] = "2"
	data["dsp_id"] = "17"
	data["adomain"] = "sportingnews.com"
	data["crid"] = "123456"
	data["media_token"] = ""
	data["media_id"] = ""
	data["media_name"] = ""
	data["bundle"] = "com.lenovo.anyshare.gps"
	data["bidfloor"] = "5000000"
	data["sh_bidfloor"] = "5000000"
	data["nurl"] = "http://baidu.com"
	data["click_url"] = "http://click.com"
	data["price"] = "5000000"
	data["sh_price"] = "5000000"
	data["sh_adm_price"] = "5000000"
	data["auction_price"] = "5000000"
	data["auction_price_raw"] = "5000000"
	return data
}
func decodeValue(str string) {
	baseStr, _ := base64.URLEncoding.DecodeString(str)
	encData, _ := encodex.UrlDecode(baseStr)
	fmt.Println(string(encData))
}
func main() {
	var arr []string
	a := &struct {
		name []string
	}{}
	fmt.Println(arr, a.name)
	for i, s := range a.name {
		fmt.Println(i, s)
	}
	return
	//fmt.Println(genData())
	//str := ",12,11"
	//println(str, len(str), "\n", str[1:], len(str[1:]))
	//return
	decodeValue("4N1OjSre-VtPrKnH7mfkcQ2VpT63MxljXceyhLYyo41mhEjMGJxgC7rcH4KReYy97nQgABIUfu4WDUBG6F1X9fvoICf4RWHj1iinJ_BfQ-_gmHoQy0wAMMsNRvJsiULR7eTedgAZs1TSTZ933LmgsUAUpL1jZNkDwO00-BCmFAj2E9g9qi4Kt6Z6eWv3pMzM7amJJRqapaANzoRs8-TWy_cK_gUYRngJVjYUYbwZxmkPA1LERu1NTh3GcfItt2Fgihmu_pAFc7KKg3f93Tovr5CZtwZaiTaaX6Z8aPptjH_xms9BBK5AasdGXzD2KEDn8xugGHgiBCq1eOvQ7kcoq_1pmIMRwN1CbX2FCwYzjaeOgdn7CiUT15ob0XX8GNf6AOo43_NzH_lFh58u22KlAl6FQTL7PnQ_uYu5eOHd0uumIDI1IC02uF1cgS-KV7EgCcjhEhCnfNpFuf3YDUyS4gGqWqRHAiue9XPxc557nIK1fuGgYigrCBE1TstuwQ6GHYtA88Ia3S-Hn64iKFzpxGonKDVTyc2PWiDbKRJa7iz44nAfNhLeeQUtKCkJXuxR69LIf4csw45oGTKOQ2diuZsbrCu2rBkYuC-9kR-RV7CbRvL9UIpKoKk58gBC9fL6SPZDn8Cy0-9TJWiFGmPPljHeaNXpgNCcXc091yI0YALleL6tD1h1nebipZZCsUb00o6FzyvAoZvUzTbDsNv6rXI9ruTTiCnhHDPQ929WA0CeoTGqCqrwevsxRreBLF_2NxxqjZohRz9OYxsvuP5-pv1bo8GtIQQQQIN6FUFQrjhscnS0bTfGyzFvEIZViCvpAhk4VdAFiMbvpGH9AZZtAFO2AvXlEuD6wIHtllCjBeZPK_6InFbnqgRzwTZBwbBuq796Epf1vnm7P0fPrR-ls2t8UfAq8I-Zdgl8EtVS3M4KdsvmSFiGVhUyNqp33o8Oq8sO6i6S6XnsKwTbuu1Q6CtPxghbJvFI6X0LS8WX25POCwFoZ-80eVvNUCN6JId9FPSnaFbZyI26ZBndc9rufgiR4hWGSzUJB_acVXuThBin59VkShpCshIo4H4M9opE-LLRTFc913k-IpNq-_XhI_OrYg2gz8VShETDx_XtLNiCIluIw9Kbuu2-Y1kIETq2Z-7T6LK0Fwu15uJ8iDFlOby8-cdKbTKdpDI3Sldon34klCGJqKHy6mnKiLi5BIRzYJYsSkQWvItQHF6FCV5p51et5vy52HoVhPO_XyKLP7Xnmbf0ML0HxEdnRuNBOg==")
	//fmt.Println(new_sh_price)
	// viewid=xthUgXjT9xh47_iVswm9f3Bb21GZZ8bgfnNzl5TMZmWZ-xolrsYlcG4F-_ECM8NGtv2iDw0h7Mqzr0rk2z5eTZ2OODWjDZwGTI-HWCLtAVgAqAk5gUoCjDehJ0JpclUzb3FCj57YQ6jiJwwOnai4hXP_hEpFYpZbjFrDZjy9pbcsoN7OqRSjTk911rJMmlfuSP1uECGz_M4_o0kUYEz388OEN4wFUI2TXA0WMFb5p7wpxLLaBvREopCHXXTxTDtWB5GTbAYPEJ3pqp_dEAUVSOkW9uOl_3Zzdfx4EWQtt3qO9vXdnqiUwVzpHIFmi-JlF9_gvvPsjJY2c5HfrGdcsUi4ymxoQI9jF9SGkxpLCYIHoh1t9srWpB7qMlfLrr-wTRMFE1nRqqXUoYtj9v2ufzLM5sYRvdONbLsgim9N7N9fIxWSWsSQ3KTb64vkPdUUAOXUkLlbTTodn9lgoWxpP6otsm0nWyGbGPPotf3VgRdUqC8VAtHiCZIKv3C5G3wGkumLyucJp2SSn29x30FqQ1StCsUF3ti0D6pLF9qJxc-7tnNAkBx1GdNF283nKFax9Mr_bHbvfy0V6BfkQVk3QAVinnkKkt79Xegp2CzEWy-PwrdNoUcOlF-OZc7gJd3le5O0z7rgS84J_QLykg3rW8jpZNAZgMCcRqPq6o0etlWxmCf-8BNcK2GiKtPsYnbzOxIp_jPxfVv1TiveBT9WYKfWeVFJOe6UP946ccxORUKT5zDgcUm9VnYnVBHAOlAePFozdgY6TeQ4acpiO-IQKAx65pAYeokx8kji3xs05NFy-uZ_ekdztYFxelybcGEzZ5nQd4gDO0fbBuhMndiaLk1IgUJQSUXz1lHgEMBrVeeTsb_QPamh_Nhk5q0QspNCF9AQN2iG6NTd2MU3-bzfFct4AisO-6WN9Lg6XfMaFSvfmO-J6yJK7icu-DQ2uWnTF9zlXYcFMu2KMiEXOihETL7uO-IG3JbOJWw401tfZScna3kZrLmlHkc9g_jvyXwJ3bTSGqBaW0Ui129LqHPg7ESAsHAa9BLG8WjCFAE0EA1MBjRrfPnejxJAx3cNuftySPbwaCI6HkncFqIdAQeMMWVt0qjIi12mL44AbQLat5P5oOFyn4IcLM67iA7mGtcYoP9dGh4vjt5sMLoPSCQH2JXuXzrZsZzVNrq4h0PvM7miBFkRsJE3jRvIfNQ2N9wE0tANHa7QrL_HcCk1_7H76fqhmB5Qti2g2RAgxWIT38_uhWomaI03w0pRnFA_ycoc3vrbha39gO06yi6_uF71AXKmBR1LVNAOiVDlW5tVciEjVuM-x_wE_f8dn1YOfjBXJ2YqfjmwbF9SD_20i-g7iKNolUhPweQdl3fHe0OL8XW8sGaB4zMgfMnIPY41y30lXZeHgpdbA_2e3Nt6dhZq7FV4DbRADJnD4w6yq2rSYYQXEg-JrTM9rjajQz1kvOLc4u3G-Sdyqlos2L1XMqlvfhynVTuxMbZwWsvQUE13t9LjeIspC41icJdXr38Au1bp6EddwZ_DO_lFEkRT6ElzT-p2PAbfZdx_aGgqqtPGE3TNMYwuxmHUFDqNgIxD8wsmHuEm4kXT4nt-u_Ac5til7SPwaYdhg__ubA7sSk4XBSK_3TKYtCWnXHjNXUPMdBJAPaR6aYG-6H2ghWYjbfQO7WFa5mmal-wlVXzzfpdDSospwkFi1y6z_CL2CoTdy98lNHWkvG00WgBrYZY6dDgrZXMO_fy8JRbx7CNmyF_wJf6QW4pFv3vQpqOvaoNlYy5c4ESgq-vlpVA3k6TB_nRc\&enc=1\&auction_price=0.031300\
	//decodeValue("4N1OjSrbqF0d_6Hb4CzmczwfpEODsFrWE5W6Rj9loKPycLMf4QLfe5ZnkY8jpQKzVvGP9yPk3v5TF1wJa3x5RKk6_arAvsvRCokI9EowqCkDPg6_h3dsQdPxRK1pREXGo3iCJ5kRG6VFyosD4uTUPwIriv_ljQs1M9j3srzWhghGP0_aXFuP4GVPjrd5tDunSCMOgEIkBun8_gFPCfgTFsmgsynRz9SZu2eXYlyghQm_Mm75LCg5zHGhzH4Uo6TIuclH_Wy3h8VLToFdwcM0pOh3D5itzQhmRvLeMof1vMbx2-y0MccO8H-qIb1jbQxF_WEWR-IeZlR-TSdjCPDpPYCDo5xZx0nNJP8nQBZjFJGrrWzw0FdftdGOhdSnVBqGbJQ83mzNGiZI_HzGHaswA1kihr738X9zV1byWs9hvDGZ8jZZvsXom8nshog9YvqyjzzDL3Um9NAWO4MXn9LzpKvYAlfqVpLA_ZDjhP9LnKOrcsgXTkpXQChs58zTxr2PYVNb-Ft20LrifmCxd9Ijfd1O8XCDMUfhOR-u9PySOfsMpYadmcFpj1c2JOH26V1rsLq53mkAVVu5m0-dv2mBxu4mG70TNrTBc6fjhOYmrhhLzbt4MdqhmH2-uTRU2n4X65BKUVNoqWlcjRkKOC8FEHY5Lvu_UTI3Qrcll0gbEzX0qzI9-a2q05iS6J2EnZfiRWGPOZ7PYpHPXzYhLr_kps9Cwq_ievK4DvYvgyQeFTkWp6UsVjJpJpRLo4tTwsPuoR1dKVhQubsmlgeij2IsxYqYI5wsEdEJuHzBC2wyskMrShFd8J1u8cK5okMqtn1dW6pysDPexSoOo7xO-yI89csy-FdouPzl7UfZuIf8BMWoQeITAOHKY--Au1BBVliut4yHDBCbcfQRB5VCcj85kh4n0mA6o9i-HZETu079Rb8PJ0K6djYqa4eNdd436mRMrAr-UQMC2OSF3HwbhhVP9hxBWmEAS825vzHm4kz6BzWXlLZm-E2xgUVy9Viun5tmUwi_gcwPgBxx_YF0t57CY1xR")
	return
	params := url.Values{}
	params.Set("bid_time", strconv.FormatInt(time.Now().Unix(), 10))
	params.Set("country", "")
	params.Set("os_id", "1")
	params.Set("osv", "2")
	params.Set("bid_id", "32f94228-52d3-4bce-be88-5dd1d799ae461233655")
	params.Set("imp_id", "")
	params.Set("udid", "3e56c2e-149d-47bb-b978-44fbe63f0a003")
	params.Set("bid_ip", "123.49.1.66")
	params.Set("ad_type", "1") // banner，video，native，native-video
	params.Set("device_type", "1")
	params.Set("network_type", "1")
	params.Set("traffic_type", "1")
	params.Set("ssp_id", "1")
	params.Set("ssp_name", "vungle")
	params.Set("adomain", "pangleglobal.com")
	params.Set("crid", "1683699610810402")
	params.Set("bundle", "com.commsource.beautyplus")
	params.Set("bidfloor", "123")
	params.Set("sh_bidfloor", "64")
	params.Set("price", "33")
	params.Set("sh_price", "77")
	params.Set("sh_adm_price", "99") // sh表示shareit，我们给DSP价格
	params.Set("dsp_id", "2")
	params.Set("dsp_name", "webeye_mock")
	params.Set("nurl", "")
	params.Set("c_size", "200*400")
	enData := params.Encode()
	encData, _ := encodex.UrlEncode([]byte(enData))
	baseStr := base64.URLEncoding.EncodeToString(encData)
	baseUrl := "http://127.0.0.1:28080/win?"
	allUrl := baseUrl + "viewid=" + string(baseStr) + "&enc=1&auction_price=${AUCTION_PRICE}"
	fmt.Println(allUrl)
}
