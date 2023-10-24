package main

import "zinx/znet"

/*
基于zinx框架开发的服务
*/

func main() {
	// 创建一个server句柄
	s := znet.NewServer("[ZINX V0.1]")
	// 启动server
	s.Server()
}
