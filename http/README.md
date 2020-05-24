`het/http`包作为Go标准库的一部分，提供了HTTP客户端和服务端的实现。使用Go来构建web服务不再需要其他的第三方包

### 简单的web server

```go
import "net/http"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":3000", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gophers!"))
}
```

`mux := http.NewServeMux()`定义了一个多路复用器(multiplexer)，用来路由进来的请求到相应的handler

### 使用Server启动web server

除了上面的那种方法启动web server，还可以使用Server来启动

```go
// from https://golang.org/src/net/http/server.go?s=77156:81268#L2480
// A Server defines parameters for running an HTTP server.

// The zero value for Server is a valid configuration.
type Server struct {
	Addr    string  // TCP address to listen on, ":http" if empty
	Handler Handler // handler to invoke, http.DefaultServeMux if nil
	TLSConfig *tls.Config
	ReadTimeout time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	MaxHeaderBytes int
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
	ConnState func(net.Conn, ConnState)
	ErrorLog *log.Logger
	BaseContext func(net.Listener) context.Context
	ConnContext func(ctx context.Context, c net.Conn) context.Context
	disableKeepAlives int32     // accessed atomically.
	inShutdown        int32     // accessed atomically (non-zero means we're in Shutdown)
	nextProtoOnce     sync.Once // guards setupHTTP2_* init
	nextProtoErr      error     // result of http2.ConfigureServer if used
	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	doneChan   chan struct{}
	onShutdown []func()
}
```

使用server重新定义一个web server

```go
import "net/http"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gopher!"))
}
```

### multiplexer

multiplexer用来定义web应用的路由


main.go
```go
import "net/http"

func main()  {
	registerRoutes()
	httpServer := http.Server{
		Addr: ":3000",
		Handler: mux,
	}

	httpServer.ListenAndServe()
}
```

routes.go

```go
import "net/http"

var mux = http.NewServeMux()

func registerRoutes()  {
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
}
```

handler.go

```go
import "net/http"

func logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout route"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login route"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about route"))

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home route"))
}
```

### handler

handler负责处理进来的http请求，在multiplexer中可以定义很多不同的路由

有2中方法来定义handler

使用struct，为了让struct是有效的handler，必须满足handler接口的定义

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

首先定义一个http handler的struct

```go
type CustomHandler struct {}
```
然后实现Handler接口
```go
func (c CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom handler!"))
}
```
使用`CustomHandler`

```go
type CustomHandler struct {}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom handler!"))
}

func main()  {
	handler := CustomHandler{}
	mux := http.NewServeMux()
	mux.Handle("/", &handler)
	http.ListenAndServe(":3000", mux)
}
```

也可以使用function作为handler，这多亏了`HandlerFunc`

```go
// source : https://golang.org/src/net/http/server.go?s=61509:61556#L1993
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

* `HandlerFunc`相当是适配器，它接受一个函数，返回一个实现了`ServeHTTP`方法的`Handler`


```go
import "net/http"

func functionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("function as http handler"))
}

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", functionHandler)
	http.ListenAndServe(":3000", mux)
}
```




