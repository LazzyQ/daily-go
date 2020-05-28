package main

import "net/http"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setHeader)
	http.ListenAndServe(":3000", mux)
}

func setHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ALLOWED",  "GET,POST")
	w.Write([]byte("set allowed headers\n"))
}
