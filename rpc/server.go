package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// 定义计算器
type Calculator struct {

}

func (c *Calculator) Multiply(args *Args, r *CalcResult) error {
	r.Result = args.Num1 * args.Num2
	return nil
}

func (c *Calculator) Sum(args *Args, r *CalcResult) error {
	r.Result = args.Num1 + args.Num2
	return nil
}


func main()  {
	calculator := new(Calculator)
	// rpc库对注册的方法有一定的限制，
	// 方法必须满足签名func (t *T) MethodName(argType T1, replyType *T2) error
	// 这个方法的第一个参数代表调用者(client)提供的参数
	// 第二个参数代表要返回给调用者的计算结果
	rpc.Register(calculator)
	// 注册 HTTP 路由
	rpc.HandleHTTP()
	// 监听端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("server error: ", err)
	}
}