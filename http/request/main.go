package main

import (
	"net/http"
	"strconv"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestInspection)
	http.ListenAndServe(":3000", mux)
}

func requestInspection(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method: " + r.Method + "\n"))
	w.Write([]byte("Protocol Version: " + r.Proto + "\n"))
	w.Write([]byte("Host: " + r.Host + "\n"))
	w.Write([]byte("Referer: " + r.Referer() + "\n"))
	w.Write([]byte("User Agent: " + r.UserAgent() + "\n"))
	w.Write([]byte("Remote Addr: " + r.RemoteAddr + "\n"))
	w.Write([]byte("Requested URL: " + r.RequestURI + "\n"))
	w.Write([]byte("Content Length: " + strconv.FormatInt(r.ContentLength, 10) + "\n"))

	for k, v := range r.URL.Query(){
		w.Write([]byte("Query string: key=" + k + " value=" + v[0] + "\n"))
	}
}
