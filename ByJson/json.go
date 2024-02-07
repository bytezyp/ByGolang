package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	var caids []map[string]interface{}
	byteStr := `[{"vendor":"0","caid":[{"caid":"00_FB649E1585AE934B3E61CA1501E6FC11_7B9CD3DDC84EBA3918C2EFF256E25135","version":"00"}],"generateTime":"1703672154"},{"vendor":"1","caid":[{"caid":"c866b8c670bf878f868c9464e7abf469","version":"20220111"},{"caid":"0583eee2559460591109c4183d680df7","version":"20230330"}],"generateTime":"1703672154"}]`
	caidJsonDecodeErr := json.Unmarshal([]byte(byteStr), &caids)
	if caidJsonDecodeErr != nil {
		fmt.Println(caidJsonDecodeErr)
		return
	}
	for _, info := range caids {
		var keyName string
		if _, isSet := info["kenyId"]; isSet { //快手的特殊格式
			keyName = "kenyId"
		}
		if _, isSet := info["qaid"]; isSet { //腾讯的特殊格式
			keyName = "qaid"
		}
		if _, isSet := info["caid"]; isSet { //其他的
			keyName = "caid"
		}
		if len(keyName) == 0 {
			continue
		}
		if _, isSet := info["vendor"]; isSet { //百度格式
			if info["vendor"] == "0" {
				continue
			}
			var subCaid []map[string]string
			fmt.Println(info["caid"])
			caidMap, err := info["caid"].([]interface{})
			fmt.Println(caidMap, err)
			//subCaidJsonDecodeErr := json.Unmarshal([]byte(), &subCaid)
			//if subCaidJsonDecodeErr != nil {
			//	continue
			//}
			for _, item := range subCaid {
				if itemCaid, unEmpty := item["caid"]; unEmpty {
					if itemVersion, isOk := item["version"]; isOk {
						fmt.Sprintf("%v_%v", itemVersion, itemCaid)
					}
				}
			}
		} else {
			//if caid := info[keyName]; caid != nil {
			//	if caidString, unEmpty := caid.(string); unEmpty {
			//		fmt.Sprintf("%v_%v", info["version"], caidString)
			//	}
			//}
		}

	}

	return
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
