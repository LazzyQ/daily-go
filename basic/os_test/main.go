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

}
