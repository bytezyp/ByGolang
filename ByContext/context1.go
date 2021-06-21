package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main1() {
	httpNum := 5
	//arr := make([]int, 5,5)
	//fmt.Println(arr)
	//for i := 0; i < httpNum; i++ {
	//	go reqContext(200, arr, i)
	//}
	//time.Sleep(1*time.Second)

	//fmt.Println(arr)
	arr := make([]int, 5, 5)
	wg := new(sync.WaitGroup)
	ctx, cancel := newContextWithTimeOut(200)
	for i := 0; i < httpNum; i++ {
		wg.Add(1)
		go reqContext2(ctx, arr, i, wg)
	}
	wg.Wait()
	fmt.Println(arr, cancel)
}
func newContextWithTimeOut(num time.Duration) (context.Context, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, num*time.Millisecond)
	return ctx, cancel
}
func reqContext2(ctx context.Context, arr []int, i int, group *sync.WaitGroup) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://baidu.com", nil)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		//fmt.Println("request err:", err)
		arr[i] = 500
		group.Done()
		return
	}
	group.Done()
	arr[i] = response.StatusCode
	return
}
func reqContext(num time.Duration, arr []int, i int) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, num*time.Millisecond)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://baidu.com", nil)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		//fmt.Println("request err:", err)
		arr[i] = 500
		return
	}
	arr[i] = response.StatusCode
	return
}
