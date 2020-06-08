package main

import (
	"log"
	"time"
)

func main()  {
	selectOnCloseChannel()
}

// 在已经close的Channel上select
// 2020/06/08 23:27:19 main start
// 2020/06/08 23:27:20 channel result:  0
// 2020/06/08 23:27:21 channel result:  0
// 2020/06/08 23:27:22 channel result:  0
// 2020/06/08 23:27:23 channel result:  0
// ...
// 从close的Channel上select会一直获取到零值
func selectOnCloseChannel()  {
	log.Println("main start")
	stopChain := make(chan int)

	close(stopChain)

	for  {
		time.Sleep(time.Second)
		select {
		case i := <- stopChain:
			log.Println("channel result: ", i)
		}
	}
}