package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// for prod use secure key, not hard-coded
var store = sessions.NewCookieStore([]byte("1234"))

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sessionHandler)
	http.ListenAndServe(":3000", mux)
}

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "custom-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	val := session.Values["hello"]
	if val != nil {
		fmt.Fprintln(w, "retrieving existing session: ")
		fmt.Fprintln(w, val)
		return
	}
	session.Values["hello"] = "world"
	session.Save(r, w)
	fmt.Fprintln(w, "no existing session found, set value for session")
}


