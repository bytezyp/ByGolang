package main

import (
	"ByGolang/report"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type websocketClientManager struct {
	conn        *websocket.Conn
	addr        *string
	path        string
	sendMsgChan chan []byte //string
	recvMsgChan chan []byte //string
	isAlive     bool
	timeout     int
}

// 构造函数
func NewWsClientManager(addrIp, addrPort, path string, timeout int) *websocketClientManager {
	addrString := addrIp + ":" + addrPort
	var sendChan = make(chan []byte, 65535)
	var recvChan = make(chan []byte, 65535)
	var conn *websocket.Conn
	return &websocketClientManager{
		addr:        &addrString,
		path:        path,
		conn:        conn,
		sendMsgChan: sendChan,
		recvMsgChan: recvChan,
		isAlive:     false,
		timeout:     timeout,
	}
}

// 链接服务端
func (wsc *websocketClientManager) dail() {
	var err error
	u := url.URL{Scheme: "ws", Host: *wsc.addr, Path: wsc.path}
	log.Printf("connecting to %s", u.String())
	wsc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return

	}
	wsc.isAlive = true
	log.Printf("connecting to %s 链接成功！！！", u.String())

}

// 发送消息
func (wsc *websocketClientManager) sendMsgThread() {
	go func() {
		for {
			msg := <-wsc.sendMsgChan
			fmt.Println("发送给服务器:", msg)
			err := wsc.conn.WriteMessage(websocket.BinaryMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				continue
			}
		}
	}()
}

// 读取消息
func (wsc *websocketClientManager) readMsgThread() {
	go func() {
		for {
			if wsc.conn != nil {
				_, message, err := wsc.conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					wsc.isAlive = false
					// 出现错误，退出读取，尝试重连
					break
				}
				log.Printf("接收到服务器发来的:%s", string(message))
				log.Printf("============================================================")
				// 需要读取数据，不然会阻塞
				wsc.recvMsgChan <- []byte(message)
			}

		}
	}()
}

// 开启服务并重连
func (wsc *websocketClientManager) start() {
	for {
		if wsc.isAlive == false {
			wsc.dail()
			wsc.sendMsgThread()
			wsc.readMsgThread()
		}
		time.Sleep(time.Second * time.Duration(wsc.timeout))
	}
}

func getSdkRequest() []byte {
	payload := make(map[string]string, 10)
	payload["appkey"] = "11111111111"
	payload["channel"] = "default"
	payload["game_status"] = "0"
	payload["interval"] = "30"
	payload["m1"] = "56f049bf15d46b8811245c7d51cf1c49"
	payload["m2"] = "c21f969b5f03d33d43e04f8f136e7682"
	payload["model"] = "HLK-AL00"
	payload["network"] = "3"
	payload["os_version"] = "10"
	payload["qid"] = ""
	payload["sdk_version"] = "772"

	requestBody := &report.ReportRequestBody{
		Product:   1,
		Type:      40,
		Version:   1,
		Timestamp: 1701834454,
		Payload:   payload,
	}

	msg, err := proto.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return msg
}

func main() {
	wsc := NewWsClientManager("180.163.249.223", "80", "/report", 10)
	payload := make(map[string]string, 10)
	payload["appkey"] = "11111111111"
	payload["channel"] = "default"
	payload["game_status"] = "0"
	payload["interval"] = "30"
	payload["m1"] = "56f049bf15d46b8811245c7d51cf1c49"
	payload["m2"] = "c21f969b5f03d33d43e04f8f136e7682"
	payload["model"] = "HLK-AL00"
	payload["network"] = "3"
	payload["os_version"] = "10"
	payload["qid"] = ""
	payload["sdk_version"] = "772"
	requestBody := &report.ReportRequestBody{
		Product:   1,
		Type:      40,
		Version:   1,
		Timestamp: 1701834454,
		Payload:   payload,
	}
	msg, err := proto.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			<-time.After(1 * time.Second)
			wsc.sendMsgChan <- msg //写入数据到通道
		}
	}()

	wsc.start()

}
