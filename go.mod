module github.com/zengqiang96/daily-go

go 1.14

require (
	github.com/apache/pulsar-client-go v0.3.0
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/sessions v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.4.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/smallnest/rpcx v0.0.0-20200924044220-f2cdd4dea15a
	github.com/sony/gobreaker v0.4.1
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/text v0.3.5
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	gonum.org/v1/gonum v0.9.1
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

// 替换为v1.26.0版本的gRPC库
replace (
	github.com/zengqiang96/dayli-go => /Users/zengqiang96/codespace/daily-go
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
