module github.com/LazzyQ/daily-go

go 1.14

require (
	github.com/apache/thrift v0.13.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/sessions v1.2.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.4.0
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/prometheus/client_golang v1.3.0
	github.com/sony/gobreaker v0.4.1
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
)

// 替换为v1.26.0版本的gRPC库
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

