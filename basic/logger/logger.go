package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

var lines = []string{
	`{"body":{"code":0,"data":{"is_sentence_match":false,"shop_question_id":"","question":""}},"custom":{"custom":{"costms":0,"status":"200 OK"}},"file":"logging.go:59","func":"middles.LoggingResponseWithLogger.func1","level":"info","msg":"outgoing http response","time":"2020-12-21T10:41:29.792+08:00","trace":"80215f3c517b70a97267a50188875120","ts":1608518489}`,
}

func main() {
	f, err := os.OpenFile("/Users/zengqiang96/xiaoduo/log/test.log", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("打开文件失败")
	}
	log.SetOutput(f)
	log.SetFlags(0)

	ticker := time.NewTicker(time.Millisecond * 500)

	for range ticker.C {
		log.Println(lines[rand.Intn(len(lines))])
	}
}
