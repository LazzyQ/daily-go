package main

import "fmt"

func main() {
	dd := make([][]int, 4)

	for i := 0; i < 4; i++ {
		dd[i] = make([]int, i+1, i+1)
	}

	for i := 0; i < len(dd); i++ {
		h := dd[i]
		for j := 0; j < len(h); j++ {
			fmt.Printf("%d ", dd[i][j])
		}
		fmt.Println()
	}

}
