package main

import (
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gopher!"))
}
