package znet

import "zinx/zinterface"

type Request struct {
	// 封装好的已经和客户端建立的链接
	conn zinterface.IConnection

	// 客户端发宋来的数据，使用msg协议封装
	//data []byte

	// v0.5 将data修改为message
	msg zinterface.IMessage
}

func (r *Request) GetConnection() zinterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
