package main

import "net/http"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", unauthorized)

	http.ListenAndServe(":3000", mux)
}

func unauthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("you do not have permission to access this resource.\n"))
}
