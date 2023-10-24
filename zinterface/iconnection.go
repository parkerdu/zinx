package zinterface

import "net"

// 定义链接的接口

type IConnection interface {
	// 启动链接
	Start()
	// 停止链接
	Stop()

	// 获取当前的链接socket
	GetTCPConnection() *net.TCPConn
	// 获取当前链接模块的id
	GetConnID() int32

	// 获取客户端的tcp状态  IP port
	RemoteAddr() net.Addr
	// 发送数据给客户端
	Send(data []byte) error
}

// 定义一个实现接口的类
type HandleFunc func(*net.TCPConn, []byte, int) error
