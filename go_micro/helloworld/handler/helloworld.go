package handler

import (
	"context"
	helloworld "github.com/LazzyQ/daily-go/go_micro/helloworld/proto/helloworld"
	log "github.com/micro/go-micro/v2/logger"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Info("Received HelloWorld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
