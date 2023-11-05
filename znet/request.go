package znet

import "zinx/zinterface"

type Request struct {
	// 封装好的已经和客户端建立的链接
	conn zinterface.IConnection

	// 客户端发宋来的数据
	data []byte
}

func (r *Request) GetConnection() zinterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
