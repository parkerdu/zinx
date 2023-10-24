package znet

import (
	"fmt"
	"net"
	"zinx/zinterface"
)

type Connection struct {
	// 当前的链接
	Conn *net.TCPConn
	// ID
	ConnID uint32

	// 当前链接状态
	isClosed bool
	// 当前链接绑定的路由处理方法
	handleAPI zinterface.HandleFunc

	// 推出chan
	ExitChan chan bool
}

// 初始化
func NewConnection(conn *net.TCPConn, connID uint32, callback zinterface.HandleFunc) *Connection {
	c := Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback,
		ExitChan:  make(chan bool),
	}
	return &c
}

func (c *Connection) StartRead() {
	fmt.Printf("read goroutine is running....")
	defer fmt.Printf("connID: %d, read is exit, remoteAddr is %s", c.ConnID, c.RemoteAddr())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			continue
		}

		if err := c.handleAPI(c.Conn, buf, n); err != nil {
			break
		}
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
func (c *Connection) GetConnID() int32 {
	return c.ConnID
}

// 获取客户端的tcp状态  IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据给客户端
func (c *Connection) Send(data []byte) error {
	n, err := c.Conn.Write(data)
	if err != nil {
		return err
	}
}
