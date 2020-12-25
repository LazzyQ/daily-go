package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var lines = []string{
	`{"app":"shop-question-server","file":"matcher.go:171","func":"question.(*SetSentenceMatcher).UpdateSentence","level":"debug","msg":"askway: 能否尽快发货, id: 5dc7ae092d7c430019abfd64","time":"2020-12-03T15:32:44.598+08:00","ts":%d, "trace": "%s"}`,
	`{"app":"shop-question-server","file":"matcher.go:171","func":"question.(*SetSentenceMatcher).UpdateSentence","level":"debug","msg":"askway: 还等什么呢?, id: 5e6f1d88a18539001d386496","time":"2020-12-03T15:32:44.597+08:00","ts":%d, "trace": "%s"}`,
	`{"app":"shop-question-server","file":"matcher.go:290","func":"question.NewWordCountIndex","level":"debug","msg":"{[过敏可以退款] [1] 5f8d24f9a53c22001854e4eb 过敏可以退款 6 2020-10-19 05:33:37.113 +0000 UTC}","time":"2020-12-03T15:32:44.596+08:00","ts":%d, "trace": "%s"}`,
}

func main() {
	f, err := os.OpenFile("/Users/zengqiang96/xiaoduo/log/test.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(f)
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		line := lines[rand.Intn(3)]
		h := strconv.FormatInt(t.Unix(), 16)
		log.Println(fmt.Sprintf(line, t.Unix(), h))
	}
}
