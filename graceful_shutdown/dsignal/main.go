package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		<-ch

		fmt.Println("func1")
	}()

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		<-ch

		fmt.Println("func2")
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
	fmt.Println("func3")
}
