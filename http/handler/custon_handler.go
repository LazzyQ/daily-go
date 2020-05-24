package main

import "net/http"

type CustomHandler struct {}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom handler!"))
}
