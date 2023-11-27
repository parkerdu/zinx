package zinterface

/*
将请求的消息封装到message中，定义一个抽象的接口
*/
type IMessage interface {
	// 获取消息属性
	GetMsgId() uint32
	GetMsgLen() int
	GetData() []byte

	SetMsgId(uint32)
	SetMsgLen(int)
	SetData([]byte)
}
