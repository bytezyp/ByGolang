package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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

type Person struct {
	Name string `json:"name,omitempty"`
}

var PwdKey = []byte("DIS**#KKKDJJSKDI")

// PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

// 实现加密
func AesEcrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 实现解密
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// 加密base64
func EnPwdCode(pwd []byte) (string, error) {
	result, err := AesEcrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

// 解密
func DePwdCode(pwd string) ([]byte, error) {
	//解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	//执行AES解密
	return AesDeCrypt(pwdByte, PwdKey)

}
func main() {
	//context.WithValue()
	//pp := Person{
	//	Name: "",
	//}

	aa := strings.Contains("11,22,33", "")
	fmt.Println(aa)
	return
	str := []byte("12fff我是的站长枯藤")
	pwd, _ := EnPwdCode(str)
	fmt.Println(pwd)
	//str := "{\"123\":123}"
	m := make(map[string]int, 2)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
	//data, _ := json.Marshal(pp)
	//fmt.Printf("%s", string(data))
	return
	funnel := make(map[string]interface{}, 1)
	funnel["aaa"] = "111"
	//funnel["bbbb"] = interface{}
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
