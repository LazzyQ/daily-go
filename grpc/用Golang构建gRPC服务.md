### 用Golang构建gRPC服务

本教程提供了Go使用gRPC的基础教程

在教程中你将会学到如何：

* 在.proto文件中定义一个服务。
* 使用protocol buffer编译器生成客户端和服务端代码。
* 使用gRPC的Go API为你的服务写一个客户端和服务器。

继续之前，请确保你已经对gRPC概念有所了解，并且熟悉protocol buffer。需要注意的是教程中的示例使用的是proto3版本的protocol buffer：你可以在Protobuf语言指南与Protobuf生成Go代码指南中了解到更多相关知识。

#### 为什么使用gRPC

我们的示例是一个简单的路线图应用，客户端可以获取路线特征信息、创建他们的路线摘要，还可以与服务器或者其他客户端交换比如交通状态更新这样的路线信息。

借助gRPC，我们可以在.proto文件中定义我们的服务，并以gRPC支持的任何语言来实现客户端和服务器，客户端和服务器又可以在从服务器到你自己的平板电脑的各种环境中运行--gRPC还会为你解决所有不同语言和环境之间通信的复杂性。我们还获得了使用protocol buffer的所有优点，包括有效的序列化（速度和体积两方面都比JSON更有效率），简单的IDL（接口定义语言）和轻松的接口更新。

#### 安装

##### 安装grpc包

首先需要安装gRPC golang版本的软件包，同时官方软件包的`examples`目录里就包含了教程中示例路线图应用的代码。

```shell script
$ go get google.golang.org/grpc
```

然后切换到`grpc-go/examples/route_guide`:目录：

```shell script
$ cd ./route_guide
```

##### 安装相关工具和插件

* 安装protocol buffer编译器

安装编译器最简单的方式是去https://github.com/protocolbuffers/protobuf/releases 下载预编译好的protoc二进制文件，仓库中可以找到每个平台对应的编译器二进制文件。这里我们以Mac Os为例，从https://github.com/protocolbuffers/protobuf/releases/download/v3.6.0/protoc-3.6.0-osx-x86_64.zip 下载并解压文件。

更新`PATH`系统变量，或者确保`protoc`放在了PATH包含的目录中了。

* 安装protoc编译器插件

```shell script
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

编译器插件`protoc-gen-go`将安装在`$GOBIN`中，默认位于`​$GOPATH/bin`。编译器`proto`c必须在`$PATH`中能找到它：

```shell script
$ export PATH=$PATH:$GOPATH/bin
```

#### 定义服务

首先第一步是使用protocol buffer定义gRPC服务还有方法的请求和响应类型，你可以在下载的示例代码[./route_guide/routeguide/route_guide.proto](./route_guide/routeguide/route_guide.proto)中看到完整的.proto文件。

要定义服务，你需要在.proto文件中指定一个具名的service

```proto
service RouteGuide {
   ...
}
```

然后在服务定义中再来定义rpc方法，指定他们的请求和响应类型。gRPC允许定义四种类型的服务方法，这四种服务方法都会应用到我们的RouteGuide服务中。

* 一个简单的RPC，客户端使用存根将请求发送到服务器，然后等待响应返回，就像普通的函数调用一样。

```proto
// 获得给定位置的特征
rpc GetFeature(Point) returns (Feature) {}
```

* 服务器端流式RPC，客户端向服务器发送请求，并获取流以读取回一系列消息。客户端从返回的流中读取，直到没有更多消息为止。如我们的示例所示，可以通过将stream关键字放在响应类型之前来指定服务器端流方法。

```proto
//获得给定Rectangle中可用的特征。结果是
//流式传输而不是立即返回
//因为矩形可能会覆盖较大的区域并包含大量特征。
rpc ListFeatures(Rectangle) returns (stream Feature) {}
```

* 客户端流式RPC，其中客户端使用gRPC提供的流写入一系列消息并将其发送到服务器。客户端写完消息后，它将等待服务器读取所有消息并返回其响应。通过将stream关键字放在请求类型之前，可以指定客户端流方法。

```proto
// 接收路线上被穿过的一系列点位, 当行程结束时
// 服务端会返回一个RouteSummary类型的消息.
rpc RecordRoute(stream Point) returns (RouteSummary) {}
```

* 双向流式RPC，双方都使用读写流发送一系列消息。这两个流是独立运行的，因此客户端和服务器可以按照自己喜欢的顺序进行读写：例如，服务器可以在写响应之前等待接收所有客户端消息，或者可以先读取消息再写入消息，或其他一些读写组合。每个流中的消息顺序都会保留。您可以通过在请求和响应之前都放置stream关键字来指定这种类型的方法。
  
```proto
//接收路线行进中发送过来的一系列RouteNotes类型的消息，同时也接收其他RouteNotes(例如：来自其他用户)
rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
```
  
我们的.proto文件中也需要所有请求和响应类型的protocol buffer消息类型定义。比如说下面的Point消息类型：

```proto
// Points被表示为E7表示形式中的经度-纬度对。
//（度数乘以10 ** 7并四舍五入为最接近的整数）。
// 纬度应在+/- 90度范围内，而经度应在
// 范围+/- 180度（含）
message Point {
  int32 latitude = 1;
  int32 longitude = 2;
}
```

#### 生成客户端和服务端代码


接下来要从我们的`.proto`服务定义生成gRPC客户端和服务端的接口。我们使用`protoc`编译器和上面安装的编译器插件来完成这些工作：

在示例proto所在文件的目录下运行：

```shell script
$ protoc -I {需要扫描的包/依赖的包} --go_out=plugins=grpc:. helloworld.proto
或
$ protoc -I {需要扫描的包/依赖的包} helloworld.proto --go_out=plugins=grpc:. 
```

运行命令后会在proto所在文件的目录下生成route_guide.pb.go文件。

pb.go文件里面包含：

* 用于填充、序列化和检索我们定义的请求和响应消息类型的所有protocol buffer代码。
* 一个客户端存根用来让客户端调用RouteGuide服务中定义的方法。
* 一个需要服务端实现的接口类型RouteGuideServer，接口类型中包含了RouteGuide服务中定义的所有方法。

