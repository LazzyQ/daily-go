package main

import (
	"fmt"
	"time"
)

func main() {
	// 初始化2s的ticker
	ticker := time.NewTicker(2 * time.Second)

	for t := range ticker.C {
		fmt.Println(t)
	}
}

func resetTimer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("now1: ", time.Now())
	// sleep一段时间，但是没有超过定时器的时间
	time.Sleep(time.Second)
	fmt.Println("now2: ", time.Now())
	//  过了1s后，定时器还没有超时，此时重置定时器的时间为2s
	if timer.Reset(2 * time.Second) {
		select {
		case t := <-timer.C:
			// now2后2s才会执行
			fmt.Println("timeout: ", t)
		}
	}
}

func normalTimer() {
	// 设置2s的定时器
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("now1: ", time.Now())
	select {
	case t := <-timer.C:
		// now1后2s后会执行
		fmt.Println("timeout: ", t)
	}
}
