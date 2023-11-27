package main

import (
	"fmt"
	"zinx/zinterface"
	"zinx/znet"
)

/*
基于zinx框架开发的服务
*/

// 为我的v0.3 服务自定义一个路由方法，目前我的服务只能注册一个路由
type myRouter struct {
	znet.BaseRoute
}

// 测试Prehandle

func (m *myRouter) PreHandle(request zinterface.IRequest) {
	fmt.Println("call preHandle")
}

// 测试Handle
func (m *myRouter) Handle(request zinterface.IRequest) {
	fmt.Println("call Handle")
	_, err := request.GetConnection().GetTCPConnection().Write(request.GetData())
	if err != nil {
		fmt.Println("handle error, ", err)
	}
	return
}

func (m *myRouter) PostHandle(request zinterface.IRequest) {
	fmt.Println("call postHandle")
}

func main() {
	// 创建一个server句柄
	s := znet.NewServer("[ZINX V0.3]")

	// v0.3 增加自定义路由
	s.AddRouter(&myRouter{})

	// 启动server
	s.Server()
}
