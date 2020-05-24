package main

import "net/http"


func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", functionHandler)
	http.ListenAndServe(":3000", mux)
}






