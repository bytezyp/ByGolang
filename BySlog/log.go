package main

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"net/url"
	"strings"
)

func GetFieldFromEventLog(logFields map[string]interface{}, fieldValue ...interface{}) interface{} {
	if len(fieldValue) == 0 {
		return ""
	}
	if _, ok := fieldValue[0].(string); !ok {
		return ""
	}
	if v, ok := logFields[fieldValue[0].(string)]; ok {
		return v
	} else {
		if len(fieldValue) > 1 {
			return fieldValue[1]
		}
		return ""
	}
}
func main() {
	//strUrl := "http%3A%2F%2Ftc.cupid.iqiyi.com%2Fdsp_conversion_cb%3Fck%3Dcc3e23fbee0b41c4b1db81c40ea19992dc1%26c%3D3e23fbee0b41c4b1db81c40ea19992dc%26d%3D66002528970%26f%3D180db566afaf4f5c%26b%3D1672300622%26pmao%3D__PAYMENT_AMOUNT__%26s%3Dfb809f08eb046197b92d57e0d7588134%26tpt%3D0"
	//
	//str, err := url.QueryUnescape(strUrl)
	//
	var allUserTableIsExist bool
	fmt.Println(allUserTableIsExist)
	//
	//str2, err := url.ParseQuery("ext_info=%3dT6H2n7u")
	//fmt.Println(str2, err)
	urlstring := "http://ocpc.baidu.com/ocpcapi/cb/actionCb?a_type={{ATYPE}}&a_value={{AVALUE}}&s=1702538018&o=1608803687000&actType=2&ext_info=AFD6PSIsImwiOiIxNzAyNTM4MDE4XzE2MDg4MDM2ODcwMDAiLCJtIjoiMTYwODgwMzY4NzAwMCIsIm4iOiIxIiwibyI6IjE4NyIsInAiOiIzIiwicSI6IjIiLCJyIjoiNCIsInMiOiI3MDAiLCJ0IjoiMTUyODQ0NDg1MjUzNiIsIngiOiIxMDAxIiwieSI6IiJ9OTNBRjFDQkYwOEE3OUU5MjhDNjEzMEZNQk9ISE9ERk0iLCJjIjoiREVFUExJTksiLCJkIjoiOTk3NDEyNjkiLCJlIjoiMzc3NzIzNDI2NyIsImYiOiIxIiwiaCI6IjE3MDI1MzgwMTgiLCJpIjoiNzgyIiwiaiI6IjE2NTEyIiwiayI6Ik5EazJOREkyTlRFeyIxNCI6IjEiLCIyMSI6IjEiLCIyNCI6IjI0IiwiMjUiOiIyIiwiMjkiOiIyIiwiMzAiOiJuSFQzckhta1BOdGtQMTB6UEhuTHJqY0wiLCIzOSI6IjAiLCI0MCI6IjAiLCI2IjoiIiwiNyI6IjAiLCI5IjoiMSIsImEiOiJDMzkxQkNCOTc3MDRBNEQ4OEVG&isMock=2&token=nHT3rHmkPNtkP10zPHnLrjcL&tokenid=nHT3rHmkPNtkP10zPHnLrjcL"
	urlStrArr := strings.Split(urlstring, "?")
	//fmt.Println(urlStrArr)
	pathMap, _ := url.ParseQuery(urlStrArr[1])
	fmt.Print(pathMap.Get("ext_info"))
	amap := make(map[string]interface{})
	amap["aa"] = "1111"
	bb := GetFieldFromEventLog(amap, "bb")
	bbb := bb.(string)
	fmt.Print(bbb, amap["aa"].(string))
	return

	//return
	// 日志级别
	defer slog.Close()
	// 日志输出

	// 日志切割
	testFile := "logDir/info.log"
	h := handler.NewBuilder().
		WithLogfile(testFile).
		WithLogLevels(slog.NormalLevels).
		WithBuffSize(8).
		WithBuffMode(handler.BuffModeBite).
		WithRotateTime(rotatefile.Every30Min).
		Build()
	slogger := slog.NewWithHandlers(h)
	slogger.Infof("ceshi %v", 4444)
	slogger.Info("ceshi %v", 11111)
}
