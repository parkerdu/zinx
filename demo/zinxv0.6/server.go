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

// 测试Handle
func (m *myRouter) Handle(request zinterface.IRequest) {
	fmt.Println("call Handle")
	// v0.3
	//_, err := request.GetConnection().GetTCPConnection().Write(request.GetData())
	//if err != nil {
	//	fmt.Println("handle error, ", err)
	//}

	// v0.5
	// 读取客户端数据
	fmt.Printf("reveive from client, msgId: %d, data: %s", request.GetMsgId(), string(request.GetData()))

	// 回写数据
	resp := []byte("这是response")
	err := request.GetConnection().Send(request.GetMsgId(), resp)
	if err != nil {
		fmt.Printf("发送数据失败")
		return
	}
	return
}

type helloRoute struct {
	znet.BaseRoute
}

// 测试Prehandle

// 测试Handle
func (h *helloRoute) Handle(request zinterface.IRequest) {
	fmt.Println("call Handle")
	// v0.3
	//_, err := request.GetConnection().GetTCPConnection().Write(request.GetData())
	//if err != nil {
	//	fmt.Println("handle error, ", err)
	//}

	// v0.5
	// 读取客户端数据
	fmt.Printf("reveive from client, msgId: %d, data: %s", request.GetMsgId(), string(request.GetData()))

	// 回写数据
	resp := []byte("这是hello, response")
	err := request.GetConnection().Send(request.GetMsgId(), resp)
	if err != nil {
		fmt.Printf("发送数据失败")
		return
	}
	return
}

func main() {
	// 创建一个server句柄
	s := znet.NewServer("[ZINX V0.6]")

	// v0.5 增加自定义路由
	s.AddRouter(0, &myRouter{})
	s.AddRouter(1, &helloRoute{})
	// 启动server
	s.Server()
}
