package zinterface

type IServer interface {
	// 启动服务
	Start()

	Stop()

	Server()
}
