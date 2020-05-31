package main

import (
	"fmt"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getCookies)
	http.ListenAndServe(":3000", mux)
}

func getCookies(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for _, cookie := range cookies{
		fmt.Fprintln(w, cookie)
	}

	cookie, err := r.Cookie("cookie-1")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, cookie)
}