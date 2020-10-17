package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main()  {
	executable, _ := os.Executable()
	fmt.Println(executable)
	fmt.Println(filepath.Split(executable))


	var m  map[string]interface{}

	var (
		a int
		ok bool
	)

	fmt.Println(m["name"])

	if a, ok = m["name"].(int);  ok {
		fmt.Println("ok")
	}
	fmt.Println(a)
}
