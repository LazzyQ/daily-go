package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name string
}

func main()  {

	wt := make(chan *Student)

	go func() {
		for  {
			select {
			case s := <- wt:
				fmt.Println(s)
			}
			time.Sleep(1 *  time.Second)
		}
	}()

	go func() {
		wt <- &Student{"xx"}
		time.Sleep(100 * time.Millisecond)
		wt <- &Student{"yy"}
		close(wt)
	}()



	time.Sleep(5 * time.Second)
}
