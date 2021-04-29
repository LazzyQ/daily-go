package main

import (
	"fmt"
	"io"
	"net/http"
	"unicode/utf8"
)

func main() {
	rsp, _ := http.Get("https://item.taobao.com/item.htm?id=615588632743")
	defer rsp.Body.Close()

	data, _ := io.ReadAll(rsp.Body)
	r, _ := utf8.DecodeRune(data)
	fmt.Println(r)
}
