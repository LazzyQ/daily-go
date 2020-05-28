package main

import (
	"fmt"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", headers)
	http.ListenAndServe(":3000", mux)
}

func headers(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	accept1 := header.Get("Accept")
	accept2 := header["Accept"]
	fmt.Fprintln(w, accept1)
	fmt.Fprintln(w, accept2)
	fmt.Fprintln(w, header)
}

