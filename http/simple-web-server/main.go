package main

import "net/http"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":3000", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gophers!"))
}
