package main

import (
	"github.com/LazzyQ/daily-go/go_micro/helloworld/handler"
	helloworld "github.com/LazzyQ/daily-go/go_micro/helloworld/proto/helloworld"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.svc.helloworld"),
		micro.Version("latest"),
		micro.Flags(
			&cli.StringFlag{
				Name: "env",
				Usage: "指定运行环境",
				Value: "dev",
				EnvVars: []string{"RUN_ENV"},
			},
		),
	)

	// Initialise service
	service.Init(
		micro.BeforeStart(func() error {
			log.Info("BeforeStart...")
			return nil
		}),
		micro.Action(func(c *cli.Context) error {
			log.Info("Action...", c.String("env"))
			return nil
		}),
	)

	// Register Handler
	helloworld.RegisterHelloWorldHandler(service.Server(), new(handler.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
