package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("%s\n", bytes)
}
