package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// normalDomain()

	wrongDomain()

	// cannotConnectIP()

	//request := gorequest.New()
	//resp, _, errs := request.Get("http://192.168.0.199").End()
	//fmt.Println("resp: ", resp, " errs: ", errs)

}

// 正常的请求，err != nil resp能够获取数据
func normalDomain() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("resp: ", resp)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body: ", string(b))
}

func wrongDomain() {
	// 不正确的域名, err != nil
	resp, err := http.Get("https://www.baidu.com1")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("resp: ", resp)
}

func cannotConnectIP() {
	// 不正确的域名, err != nil
	resp, err := http.Get("http://192.168.0.199")
	//resp, err := http.Get("http://127.0.0.1:3000")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("resp: ", resp)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body: ", string(b))
}
