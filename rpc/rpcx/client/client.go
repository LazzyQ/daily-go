package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/zengqiang96/daily-go/rpc/rpcx/proto"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith1", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &proto.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &proto.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}

}
