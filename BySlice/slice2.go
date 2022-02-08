package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main2() {
	//s := make([]int, 0, 1)
	//s2 := make([]int,  3)
	//fmt.Println(s,s2)
	//fmt.Println(s, len(s), cap(s))
	//fmt.Printf("s: len:%d , cap:%d \n", len(s), cap(s))
	//s = append(s, 1, 2, 3)
	//fmt.Printf("append s: len:%d , cap:%d \n", len(s), cap(s))
	//s1 := append(s, 4, 5, 6)
	//fmt.Printf("append s1: len:%d , cap:%d", len(s1), cap(s1))
	log.SetFormatter(&log.JSONFormatter{})
	req, _ := http.NewRequest("POST", "127.0.0.1", nil)
	println(req)
	println(log.Fields{"req": req})
	log.WithFields(log.Fields{"req": req}).Error("new request failed")

	//var adResps []*string
	//
	//println(adResps)
	//log.WithFields(log.Fields{ "req": adResps, }).Error("new request failed2")

}
