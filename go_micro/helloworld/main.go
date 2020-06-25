package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"helloworld/handler"

	helloworld "helloworld/proto/helloworld"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.svc.helloworld"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()


	// Register Handler
	helloworld.RegisterHelloWorldHandler(service.Server(), new(handler.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
