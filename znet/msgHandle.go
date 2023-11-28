package znet

import (
	"fmt"
	"zinx/zinterface"
)

type MsgHandle struct {
	//
	method map[uint32]zinterface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{method: make(map[uint32]zinterface.IRouter)}
}

// 调度、执行对应的路由
func (m *MsgHandle) DoMsgHandle(request zinterface.IRequest) {
	// step1: 获取msgID
	msgId := request.GetMsgId()
	// step2: 根据msgID进行req处理
	handle, ok := m.method[msgId]
	if !ok {
		fmt.Printf("msgId %d do not found method", msgId)
	}
	handle.PreHandle(request)
	handle.Handle(request)
	handle.PostHandle(request)
}

// 为某个消息添加具体处理逻辑
func (m *MsgHandle) AddRoute(msgId uint32, router zinterface.IRouter) {
	// 是否存在
	if _, ok := m.method[msgId]; ok {
		panic(fmt.Sprintf("msgId: %d already exist", msgId))
	}
	m.method[msgId] = router
}
