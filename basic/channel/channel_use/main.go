package main

import (
	"log"
	"time"
)

func main()  {
	log.Println("main start")
	stopChain := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		stopChain <- struct{}{}
	}()

	select {
	case <- stopChain:
		log.Println("main end")
	}
}



