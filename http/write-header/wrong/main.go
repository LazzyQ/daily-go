package main

import "net/http"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setHeader)
	http.ListenAndServe(":3000", mux)
}

func setHeader(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bad request!\n")) // 这将会设置status为 200 ok
	w.WriteHeader(http.StatusBadRequest) // 这里的WriteHeader设置将无效
}
