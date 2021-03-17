package basic

import (
	"testing"
	"time"
)

// 在已经close的Channel上select
//
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  1
// daily-go/basic/channel/channel_test.go:27: channel result:  2
// daily-go/basic/channel/channel_test.go:27: channel result:  3
// daily-go/basic/channel/channel_test.go:27: channel result:  4
// daily-go/basic/channel/channel_test.go:27: channel result:  5
// daily-go/basic/channel/channel_test.go:27: channel result:  6
// daily-go/basic/channel/channel_test.go:27: channel result:  7
// daily-go/basic/channel/channel_test.go:27: channel result:  8
// daily-go/basic/channel/channel_test.go:27: channel result:  9
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// daily-go/basic/channel/channel_test.go:27: channel result:  0
// ...
// 从close的Channel上select会一直获取到零值
func TestSelectOnCloseChannel(t *testing.T) {
	stopChain := make(chan int, 10)
	for i := 0; i < 10; i++ {
		stopChain <- i
	}
	close(stopChain)

	for {
		time.Sleep(time.Second)
		select {
		case i := <-stopChain:
			t.Log("channel result: ", i)
		}
	}
}

// daily-go/basic/channel/channel_test.go:38: 0
// daily-go/basic/channel/channel_test.go:38: 1
// daily-go/basic/channel/channel_test.go:38: 2
// daily-go/basic/channel/channel_test.go:38: 3
// daily-go/basic/channel/channel_test.go:38: 4
// daily-go/basic/channel/channel_test.go:38: 5
// daily-go/basic/channel/channel_test.go:38: 6
// daily-go/basic/channel/channel_test.go:38: 7
// daily-go/basic/channel/channel_test.go:38: 8
// daily-go/basic/channel/channel_test.go:38: 9
//
// 在已经close的channel上for range会在channel close后退出循环
func TestRangeOnCloseChannel(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for i := range ch {
		t.Log(i)
	}
}
