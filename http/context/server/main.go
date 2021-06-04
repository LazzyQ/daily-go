package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe("", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	done := make(chan int, 1)

	go func() {
		time.Sleep(2 * time.Second)
		done <- 1
	}()

	select {
	case <-r.Context().Done():
		fmt.Println("超时了", time.Since(start).Milliseconds())
	case <-done:
		fmt.Println("正常结束了")
	}
	w.Write([]byte("index"))
}
