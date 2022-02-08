package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
	"strings"
)

var (
	client *redis.Client
)

func init() {

	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

}

func batchSet() {

	pipeline := client.Pipeline()

	//ctx := context.Background()

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("%d", i)
		pipeline.HSet(key, map[string]interface{}{key: key})

	}

	_, err := pipeline.Exec()

	if err != nil {

		panic(err)

	}

}

func batchGet() {

	//pipeline := client.Pipeline()
	//
	//ctx := context.Background()
	//
	//result := make([]*redis.StringStringMapCmd, 0)
	//
	//for i := 0; i < 100; i++ {
	//
	//	key := fmt.Sprintf("%d", i)
	//
	//	result = append(result, pipeline.HGetAll(ctx, key))
	//
	//}
	//
	//_, _ = pipeline.Exec(ctx)
	//
	//for _, r := range result {
	//
	//	v, err := r.Result()
	//
	//	if err != nil {
	//
	//		panic(err)
	//
	//		//fmt.Println(err)
	//
	//	}
	//
	//	fmt.Println(v)
	//
	//}

}

func main() {
	//pip := client.Pipeline()
	//cmds := make([]*redis.BoolCmd, 2)
	//cmds[0] = pip.HSet("key", "qtt")
	//cmds[1] = pip.Expire("key", 600)
	//_, err = pip.Exec()

	mm := make(map[string]interface{}, 2)
	mm["11"] = "1111"
	for _, i := range mm {
		if str, ok := i.(string); ok {
			fmt.Println("%v", str)
		}
	}
	freqValue := "1_123"
	freqSli := strings.Split(freqValue, "_")
	stampTime, err := strconv.ParseInt(freqSli[1], 10, 64)
	freq, err1 := strconv.ParseInt(freqSli[0], 10, 64)
	if err != nil || err1 != nil {
		return
	}
	freq++

	fmt.Println(freq, stampTime)
	freqStr := strconv.Itoa(int(freq))
	fmt.Println(freqStr)
	fmt.Printf("%T", freqStr)
}
