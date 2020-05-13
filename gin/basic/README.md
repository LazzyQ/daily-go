### gin基本使用

首先使用`gin.Default()`初始化一个默认配置的`Engine`实例，这个实例默认使用了`Logger`和`Recovery`的中间件

通过这个`Engine`实例配置路由与处理器handler

```go
r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
})

r.GET("/user/:name", func(c *gin.Context) {
    user := c.Params.ByName("name")
    value, ok := db[user]
    if ok {
        c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
    } else {
        c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
    }
})
```

当问这个2个路径就能看到相应的结果

```text
curl localhost:8080/ping
pong

curl localhost:8080/user/foo
{"status":"no value","user":"foo"}
```

通过这个`Engine`实例创建一个路由组，并使用`BasicAuth`的中间件，也就是说如果需要访问这个路由组需要提供 认证信息，也就是用户名和密码

```go
authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
    "foo": "bar", // user:foo password:bar
    "manu": "123", // user:manu password:123
}))
```

然后配置这个需要认证的路由组具体的路径与处理器handler

```go
authorized.POST("admin", func(c *gin.Context) {
    user := c.MustGet(gin.AuthUserKey).(string)
    var json struct{
        User string `json:"user"`
        Value string `json:"value" binding:"required"`
    }

    if c.Bind(&json) == nil {
        db[user] = json.Value
        c.JSON(http.StatusOK, gin.H{"status": "ok", "json": json})
    }
})
```

通过`--user`指定用户名和密码请求这个需要认证的接口

```text
curl --user foo:bar -X POST -H "Content-Type:application/json" -d '{"value": "bar"}' localhost:8080/admin

{"json":{"user":"","value":"bar"},"status":"ok"}
```

最后再试试前面的`/user/:name`，现在应该可以获取到了

```text
curl localhost:8080/user/foo
{"user":"foo","value":"bar"}
```

