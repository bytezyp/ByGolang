package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	coding_key = "dataupsdk12798ak"
	iv_key     = "dataupsdk12798ak"
)

func AesEncrypt(plaintext []byte) ([]byte, error) {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(coding_key))
	if err != nil {
		log.Printf("Error: NewCipher(%d bytes) = %s", len(coding_key), err)
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(c, []byte(iv_key))
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

func AesDecrypt(ciphertext []byte) ([]byte, error) {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(coding_key))
	if err != nil {
		log.Printf("Error: NewCipher(%d bytes) = %s", len(coding_key), err)
		return nil, err
	}

	cfbdec := cipher.NewCFBDecrypter(c, []byte(iv_key))
	plaintextCopy := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)

	return plaintextCopy, nil
}
func AesECBEncrypt(ciphertext []byte) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(coding_key))
	if err != nil {
		log.Printf("Error: NewCipher(%d bytes) = %s", len(coding_key), err)
		return nil, err
	}
	length := (len(ciphertext) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, ciphertext)
	pad := byte(len(plain) - len(ciphertext))
	for i := len(ciphertext); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(ciphertext); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted, nil
}
func AesECBDecrypt(ciphertext []byte) (decrypted []byte, err error) {
	cipher, err := aes.NewCipher([]byte(coding_key))
	if err != nil {
		log.Printf("Error: NewCipher(%d bytes) = %s", len(coding_key), err)
		return nil, err
	}
	decrypted = make([]byte, len(ciphertext))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(ciphertext); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], ciphertext[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}
func httpPostForm() {

	//clickYear, clickMonth, clickDay := time.UnixMilli(int64(timeFloat64)).Date()
	//clickDate := clickYear*10000 + int(clickMonth)*100 + clickDay
	resp, err := http.PostForm("http://112.64.200.122/heartbeat/web", url.Values{
		"product":     {"3"},
		"type":        {"40"},
		"version":     {"1"},
		"timestamp":   {"1703067324"},
		"appkey":      {"wjcq"},
		"game_status": {"0"},
		"interval":    {"120"},
		"qid":         {"122796609311111"},
	})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
func postForm() {
	// 提交数据 key-value
	//data := map[string]interface{}{
	//	"sdk_event_id": "sdk_recharge",
	//	"os":           "android",
	//	"test":         "0",
	//	"dtime":        "2024-01-04 16:29:35",
	//	"msec":         "1704356975026",
	//	"appid":        "21e0d480623a4a0b9d8074cb2dfb9756",
	//	"android_id":   "6470cfa961d98de2",
	//	"gamechannel":  "rh_215",
	//	"oaid":         "09963a537a4932c542e795dadeb5e20e94718d671567db78522407b1619709be",
	//	"qid":          "222222",
	//	"server_id":    "bbbbbb",
	//	"role_id":      "55555",
	//	"orderamount":  "200",
	//	"sn":           "5555555555555",
	//	"payid":        "2222",
	//}
	//jsonByte, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	jsonByte := []byte("{\"oaid\":\"b8c1c494335318f\",\"android_id\":\"9b853160df1ac71a\",\"msec\":1705549034038,\"dtime\":\"2024-01-18 11:37:14\",\"appid\":\"\",\"gamechannel\":\"rh_0\",\"os\":\"android\",\"test\":0,\"sdk_event_id\":\"sdk_activate\"}")
	mm := make(map[string]interface{})
	err := json.Unmarshal(jsonByte, &mm)
	if err != nil {
		fmt.Println(err)
		return
	}
	//加密 返回加密string
	encByteArr, _ := AesECBEncrypt(jsonByte)
	encStr := base64.StdEncoding.EncodeToString(encByteArr)
	resp, err := http.PostForm("https://data-up.mgame.360.cn/v1/up", url.Values{
		//"enc_key": {"gBHt127p5t9K7coNdjhQDXmnkuR7a9kglilh70aWstYo%2FMaddEOXio%2F2BdOzHYRnZj2QqIROQwNn%0AEP9ECzpSlscB3zmVtPaJolTa5lzczUQp3q%2FTG1qDb5bUmGu9%2BuiIWamFvNSmcu43ksGc%2FzTgOGDe%0AAO5Pa6TN%2BInN%2Fwh%2B1tD%2BICDjbjNSjx87pCR9HuIg4L4ZkSGCzPLQKJaNQxjylQUUr%2BoF2ypRLIUX%0A14ndnL5MlcMYkTxgKUx16mFmISpg8srB%0A"},
		"enc_key": {encStr},
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(111, string(body))
}
func main() {
	postForm()
	return
	//str2 := "gBHt127p5t9K7coNdjhQDXmnkuR7a9kglilh70aWstYo/MaddEOXio/2BdOzHYRnZj2QqIROQwNnEP9ECzpSlscB3zmVtfqFpVDU5FrWzUSdK6tngZoR8TLtsGGLliDpgXijKlhnN7RprmWVpzTY3/Vn8Ub5DbA3OLa7hyxMOrjmj7aTnxy9AiXcYgFNS8IOF7Qtn13BM/Mb8kvVRXKP9gFW9PjrvFN7lHRgoOwb8pqtIbsxXYBeBJL1XCxaxG05P6FU"
	//byteArr, _ := AesDecrypt([]byte(str2))
	//mm := make(map[string]string)
	//dd := json.Unmarshal(byteArr, &mm)
	//fmt.Println(string(byteArr), dd, mm)
	//postForm()
	//strMap := map[string]string{
	//	"aa": "111",
	//	"bb": "222",
	//}
	//fmt.Println(strMap["ccc"], len(strMap["ccc"]))
	aa := "123456"
	encStr, _ := AesECBEncrypt([]byte(aa))
	base64EncodStr := base64.StdEncoding.EncodeToString(encStr)
	base64DeocdeStr, _ := base64.StdEncoding.DecodeString(base64EncodStr)
	dercStr, _ := AesECBDecrypt(base64DeocdeStr)
	fmt.Println(string(encStr), base64EncodStr, string(base64DeocdeStr), string(dercStr))
	return
	//dd := `20220111_5d187200197671843b93adcbe4bc9d9b,20230330_6f48b827ccf01f67fa142e773df7ab61`
	//ok, matchErr := regexp.Match("^[0-9]+_[0-9a-z]+(,[0-9]+_[0-9a-z]+)?$", []byte(dd))
	//fmt.Println(ok, matchErr)
	//if matchErr == nil && ok {
	//	caidStrs := strings.Split(dd, ",")
	//	for _, caidStr := range caidStrs {
	//		fmt.Println(caidStr)
	//	}
	//}
	//return
	a := `http%3A%2F%2Ftc.cupid.iqiyi.com%2Fdsp_conversion_cb%3Fck%3Dcc3e23fbee0b41c4b1db81c40ea19992dc1%26c%3D3e23fbee0b41c4b1db81c40ea19992dc%26d%3D66002528970%26f%3D180db566afaf4f5c%26b%3D1672300622`
	a = `http%3A%2F%2Ftc.cupid.iqiyi.com%2Fdsp_conversion_cb%3Fck%3Dcc5b03fc1fee7884c6c992f22f907c21b31%26c%3D5b03fc1fee7884c6c992f22f907c21b3%26d%3D73000117090%26f%3D78bb5c68d8fdb53f3448df2988292a0adbae5522%26b%3D1704798113%26pmao%3D__PAYMENT_AMOUNT__%26cst%3D__CONVERSION_TIME__%26s%3D6a4f03fd28180bb60fc64d7a78d6212c%26tpt%3D0%26iigc%3D1%26ai%3D51200215860`
	newUrl, _ := url.QueryUnescape(a)
	strArr := strings.Split(newUrl, "?")
	vv, _ := url.ParseQuery(strArr[1])
	fmt.Println(newUrl, vv)
	str, _ := url.ParseQuery(`20220111_5d187200197671843b93adcbe4bc9d9b%2C20230330_6f48b827ccf01f67fa142e773df7ab61`)
	fmt.Println(str)
	return
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted)
	return
	//loginDayTime := time.Now()
	//loginYear, loginMonth, loginDay := time.UnixMilli(loginDayTime.UnixMilli()).AddDate(0, 0, -7).Format()
	//login7DaysAgoDate := loginYear*10000 + int(loginMonth)*100 + loginDay
	//fmt.Println(login7DaysAgoDate, loginYear, loginMonth, loginDay)
	//return
	for true {
		<-time.After(1 * time.Second)
		httpPostForm()
	}
}
