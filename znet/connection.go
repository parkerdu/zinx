package znet

import (
	"fmt"
	"net"
	"zinx/zinterface"
)

// todo 修改成session
type Connection struct {
	// 当前的链接
	Conn *net.TCPConn
	// ID
	ConnID uint32

	// 当前链接状态
	isClosed bool

	// 推出chan
	ExitChan chan bool

	// v0.3 增加router字段
	Router zinterface.IRouter
}

// 初始化, 每个conn --> 对应一个route处理方法
func NewConnection(conn *net.TCPConn, connID uint32, route zinterface.IRouter) *Connection {
	c := Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool),
		Router:   route,
	}
	return &c
}

func (c *Connection) StartRead() {
	fmt.Printf("read goroutine is running....")
	defer fmt.Printf("connID: %d, read is exit, remoteAddr is %s", c.ConnID, c.RemoteAddr())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			continue
		}

		//// v0.2 版本直接调用handleAPI, 在v0.3中替换为调用该链接对应的理由方法
		//if err := c.handleAPI(c.Conn, buf, n); err != nil {
		//	break
		//}

		// v0.3 改成调用该链接对应的路由来处理消息
		req := Request{
			conn: c,
			data: buf,
		}
		c.Router.PreHandle(&req)
		c.Router.Handle(&req)
		c.Router.PostHandle(&req)
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start, connId=", c.ConnID)

	//从链接中读取数据, 并处理业务
	go c.StartRead()
}

// 停止链接
func (c *Connection) Stop() {
	fmt.Println("conn stop, connID=", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	// 关闭socket链接
	c.Conn.Close()
	close(c.ExitChan)
}

// 获取当前的链接socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接模块的id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取客户端的tcp状态  IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据给客户端
func (c *Connection) Send(data []byte) error {
	_, err := c.Conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}
