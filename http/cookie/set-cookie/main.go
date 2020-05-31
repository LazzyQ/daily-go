package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setCookies)
	http.ListenAndServe(":3000", mux)
}

func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "cookie-1",
		Value: "hello world",
	}
	http.SetCookie(w, &cookie)
}
