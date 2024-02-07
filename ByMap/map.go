package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
	m := make(map[string]struct{}, 2)
	m["11"] = struct{}{}
	m["22"] = struct{}{}
	for s, _ := range m {
		fmt.Printf("key:%s ", s)
	}
	fmt.Printf("%v", m)
	return
	//m1 := make(map[int]int)
	//fmt.Println(m1)
	//m1[1] = 11
	//m1[2] = 12
	//m1[4] = 14
	//m1[3] = 13
	//
	//fmt.Println(len(m1))
	//for v, k := range m1 {
	//	fmt.Println(v, k)
	//}
	//m2 := make(map[string]map[string]int, 10)
	//m := make(map[string]int, 1)
	//m["111"] = 111
	//m2["a"] = m
	//m3 := make(map[string]int, 1)
	//m3["222"] = 222
	//m2["b"] = m3
	//for s, m4 := range m2 {
	//	fmt.Println(s, m4)
	//}
	//m2 := make(map[string][]int, 1470000)
	//arr := make([]int,0,10)
	//fmt.Printf("slice size: %d\n", unsafe.Sizeof(arr))
	//for i := 0; i < 100000; i++ {
	//	arr = append(arr, 10296005)
	//}
	////fmt.Println(arr)
	//fmt.Printf("slice size: %d\n", unsafe.Sizeof(arr))
	//fmt.Println("Size of [200]int:", unsafe.Sizeof([200]int{}))
	//fmt.Println("Size of [180]int:", unsafe.Sizeof([180]int{}))
	//m2["b95fb94c892f4051ab3905033ff48574"] = arr
	//m2["b95fb94c892f4051ab3905033ff48571"] = arr
	//fmt.Println(m2)
	//fmt.Printf("map size: %d\n", unsafe.Sizeof(m2["b95fb94c892f4051ab3905033ff48571"]))
	//fmt.Printf("map size: %d\n", unsafe.Sizeof(m2["b95fb94c892f4051ab3905033ff48574"]))
	//smap := make(map[int64]struct{})
	//var ids []int64
	//for id, _ := range smap {
	//	ids = append(ids, id)
	//}
	//sort.Sort(Int64Slice(ids))
	//fmt.Println(ids)
	//fmt.Printf("%T \n",ids)
	//fmt.Printf("%V \n",ids)
	//aa := make([]int64, 0, 0)
	//fmt.Printf("%T \n",aa)
	//fmt.Printf("%V \n",aa)

	//fmt.Println(len("ac49a2567bab409f99e01b7d84b64036"))
	//md5hash := md5.New()
	//md5hash.Write([]byte("06960aae437ed3c8"))
	//oaidMd5Str := hex.EncodeToString(md5hash.Sum(nil))
	//fmt.Println(oaidMd5Str, len(oaidMd5Str))
	type LogInfo struct {
		ReturnedIds       []int64 `json:"returned_ids,omitempty"`
		ReturnedIdsString string  `json:"returned_ids,omitempty"`
	}
	dd := make([]int64, 1)
	dd[0] = 10157191
	var ids []string
	for _, value := range dd {
		newId := strconv.Itoa(int(value))
		//fmt.Println(newId,value)
		ids = append(ids, newId)
	}
	fmt.Println(ids)
	//var log LogInfo
	//log.ReturnedIds = dd
	js, _ := json.Marshal(ids)
	fmt.Println(string(js))
	var personFromJSON interface{}
	json.Unmarshal(js, &personFromJSON)
	//db := json.Unmarshal(js)
	//fmt.Println(dd)
	//fmt.Println(personFromJSON)
}
