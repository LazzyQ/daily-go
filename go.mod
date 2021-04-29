module github.com/zengqiang96/daily-go

go 1.14

require (
	github.com/apache/pulsar-client-go v0.3.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/fatih/structs v1.1.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/sessions v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/miekg/dns v1.1.27 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/smallnest/rpcx v1.6.2
	github.com/sony/gobreaker v0.4.1
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/urfave/cli/v2 v2.2.0
	github.com/xtaci/lossyconn v0.0.0-20200209145036-adba10fffc37 // indirect
	golang.org/x/sys v0.0.0-20210304124612-50617c2ba197 // indirect
	golang.org/x/text v0.3.5
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.33.1
	google.golang.org/grpc/examples v0.0.0-20210429011145-91d8f0c916d7 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/ini.v1 v1.44.0 // indirect
)

// 替换为v1.26.0版本的gRPC库
replace github.com/zengqiang96/dayli-go => /Users/zengqiang96/codespace/daily-go
