package main

import "log"

func main() {
	func1()
}

func func1() int {
	log.Print("func1 start...")
	defer log.Print("func1 end...")

	{
		log.Print("static start...")
		defer log.Print("static end...")
	}

	return func2()
}

func func2() int {
	log.Print("func2 start...")
	defer log.Print("func2 end...")
	return 1
}
