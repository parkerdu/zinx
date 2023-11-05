package zinterface

/*
路由 -- 处理Irequest
*/

type IRouter interface {
	PreHandle(request IRequest)

	Handle(request IRequest)

	PostHandle(request IRequest)
}
