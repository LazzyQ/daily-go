package basic

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second)

	// for tm := range timer.C {
	// 	t.Log(tm)
	// 	timer.Reset(time.Second)
	// }

	var ch chan int
	for {
		select {
		case tm := <-timer.C:
			t.Log(tm)
			// 手动重置时间，与ticker是有却别的
			timer.Reset(time.Second)
		case <-ch:
		}
	}
}

func TestAfter(t *testing.T) {
	var ch chan int
	select {
	case tm := <-time.After(time.Second):
		t.Log(tm)
	case <-ch:
	}
}

func TestAfterFunc(t *testing.T) {
	var ch chan int
	timer := time.AfterFunc(time.Second, func() {
		t.Log("我执行了")
		ch <- 0
	})
	defer timer.Stop()
	<-ch
}

func TestTick(t *testing.T) {
	for tm := range time.Tick(time.Second) {
		t.Log(tm)
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	var ch chan int
	for {
		select {
		case tm := <-ticker.C:
			t.Log(tm)
		case <-ch:
		}
	}
}
