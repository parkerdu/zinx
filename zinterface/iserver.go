package zinterface

type IServer interface {
	// 启动服务
	Start()

	Stop()

	Server()

	// v0.3 添加路由功能, 给当前服务注册一个路由方法
	AddRouter(msgId uint32, route IRouter)
}
