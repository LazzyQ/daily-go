package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	helloworld "github.com/zengqiang96/daily-go/go_micro/helloworld/proto/helloworld"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Info("Received HelloWorld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
