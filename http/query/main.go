package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", showQuery)
	http.ListenAndServe(":3000", mux)
}

func showQuery(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	w.Write([]byte("query strings are\n"))
	w.Write([]byte("Name:" + queryString.Get("name") + "\n"))
	w.Write([]byte("Email:" + queryString.Get("email") + "\n"))
}
