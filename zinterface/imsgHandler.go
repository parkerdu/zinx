package zinterface

/*
消息管理抽象层
*/

type MsgHandler interface {
	// 调度、执行对应的路由
	DoMsgHandle(request IRequest)

	// 为某个消息添加具体处理逻辑
	AddRoute(msgId uint32, router IRouter)
}
