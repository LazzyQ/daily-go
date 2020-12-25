package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/zengqiang96/daily-go/rpc/rpcx/proto"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	// s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(proto.Arith1), "")
	s.Serve("tcp", *addr)
}
