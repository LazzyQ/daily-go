package main

import (
	"log"
	"net/rpc"
)

func main()  {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	args := &Args{
		7, 8,
	}

	var r CalcResult
	err = client.Call("Calculator.Sum", args, &r)
	if err != nil {
		log.Fatal("Sum error: ", err)
	}
	log.Printf("Sum %d + %d = %d\n", args.Num1, args.Num2, r.Result)

	err = client.Call("Calculator.Multiply", args, &r)
	if err != nil {
		log.Fatal("Multiply error: ", err)
	}
	log.Printf("Multiply %d * %d = %d\n", args.Num1, args.Num2, r.Result)
}