package main

import (
	"fmt"
	"runtime/debug"
)

//go:noinline
func sliceParam(nums []int) {
	l := len(nums)
	fmt.Println(l)
	fmt.Println(string(debug.Stack()))
}

func main() {
	var x = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		x = append(x, i)
	}
	sliceParam(x)
}
