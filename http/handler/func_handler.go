package main

import "net/http"

func functionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("function as http handler"))
}
