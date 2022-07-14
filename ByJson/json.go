package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func GetStringArrMd5(strArr []string) string {
	md5hash := md5.New()
	for _, str := range strArr {
		md5hash.Write([]byte(str))
	}
	return hex.EncodeToString(md5hash.Sum(nil))
}

func main() {
	var data = []byte(`{"status":200}`)
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	//fmt.Printf("%t",result)

	strArr := []string{"123", "1635488484", "jifen"}
	md5Str := GetStringArrMd5(strArr)
	println(md5Str)
}

//func Test_huawei_proto(t *testing.T) {
//	Convey("PackLogReqData", t, func() {
//		req := new(HuaWeiRtaReq)
//		str := `{"req_id":21231211212,"oaids":["xasxd****djhklvbn","ikbj****nvbn"],"req_time":1600000000,"campaign_id":"133****0010"}`
//		err := json.Unmarshal([]byte(str), req)
//		fmt.Println(err, req)
//		So(err, ShouldEqual, nil)
//		So(req.ReqId, ShouldEqual, 21231211212)
//
//
//		resp := new(HuaWeiRtaResponse)
//		str = `{"oaids":["xasxd****djhklvbn","ikbj****nvbn"],"rsp_time":1620000000000,"materials":[{"title":"test title","summary":"test summary","body":"test body","image_id":"133****021","image_url":"https://lfconte*****816160828.png","action_type":2,"data":"{\"param1\":\"this is param1\",\"param2\":\"this isparam2\"}"}],"result_code":"0","result_message":"success"}`
//		err = json.Unmarshal([]byte(str), resp)
//		So(err, ShouldEqual, nil)
//		So(resp.RspTime, ShouldEqual, 1620000000000)
//		So(resp.Oaids[0], ShouldNotBeNil)
//		So(resp.Oaids[0], ShouldEqual, "xasxd****djhklvbn")
//		So(resp.Materials[0], ShouldNotBeNil)
//		So(resp.Materials[0].Data, ShouldEqual, `{"param1":"this is param1","param2":"this isparam2"}`)
//	})
//}
