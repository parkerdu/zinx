package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/zinterface"
)

// todo 修改成session
type Session struct {
	// 当前的链接
	Conn *net.TCPConn
	// ID
	ConnID uint32

	// 当前链接状态
	isClosed bool

	// 推出chan
	ExitChan chan bool

	// v0.3 增加router字段
	//Router zinterface.IRouter

	// v0.6 修改router为handle，多路由
	Handle zinterface.MsgHandler
}

// 初始化, 每个conn --> 对应一个route处理方法
func NewSession(conn *net.TCPConn, connID uint32, handle zinterface.MsgHandler) *Session {
	c := Session{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool),
		Handle:   handle,
	}
	return &c
}

// 收包
func (s *Session) Receive() {
	fmt.Printf("read goroutine is running....")
	defer fmt.Printf("connID: %d, read is exit, remoteAddr is %s", s.ConnID, s.RemoteAddr())
	defer s.Stop()

	for {
		//// v0.2 版本直接调用handleAPI, 在v0.3中替换为调用该链接对应的理由方法
		//buf := make([]byte, config.MaxPackageSize())
		//_, err := s.Conn.Read(buf)
		//if err != nil {
		//	continue
		//}
		//if err := s.handleAPI(s.Conn, buf, n); err != nil {
		//	break
		//}

		// v0.3 改成调用该链接对应的路由来处理消息
		//buf := make([]byte, config.MaxPackageSize())
		//_, err := s.Conn.Read(buf)
		//if err != nil {
		//	continue
		//}
		//req := Request{
		//	conn: s,
		//	data: buf,
		//}

		// v0.5 改成安装协议拆msg包
		// step1: 创建一个拆包对象
		dp := NewDataPack()
		// step2: 读取客户端发送的数据头
		header := make([]byte, dp.GetHeaderLen())
		offset := 0
		for offset < dp.GetHeaderLen() {
			n, err := s.Conn.Read(header[offset:])
			if err != nil {
				return
			}
			offset += n
		}

		// step3: 根据数据头解包
		imsg, err := dp.UnPack(header)
		if err != nil {
			return
		}

		// step3: 根据msg.datalen得到实际数据data
		offset = 0
		buf := make([]byte, imsg.GetMsgLen())
		for offset < imsg.GetMsgLen() {
			n, err := s.Conn.Read(buf)
			if err != nil {
				return
			}
			offset += n
		}
		imsg.SetData(buf)
		// step4: 将msg存储到req中
		req := Request{
			conn: s,
			msg:  imsg,
		}
		// 使用handle 处理request消息
		go s.Handle.DoMsgHandle(&req)
	}
}

func (s *Session) Start() {
	fmt.Println("conn start, connId=", s.ConnID)

	//从链接中读取数据, 并处理业务
	go s.Receive()
}

// 停止链接
func (s *Session) Stop() {
	fmt.Println("conn stop, connID=", s.ConnID)
	if s.isClosed {
		return
	}
	s.isClosed = true
	// 关闭socket链接
	s.Conn.Close()
	close(s.ExitChan)
}

// 获取当前的链接socket
func (s *Session) GetTCPConnection() *net.TCPConn {
	return s.Conn
}

// 获取当前链接模块的id
func (s *Session) GetConnID() uint32 {
	return s.ConnID
}

// 获取客户端的tcp状态  IP port
func (s *Session) RemoteAddr() net.Addr {
	return s.Conn.RemoteAddr()
}

// 发送数据给客户端
func (s *Session) Send(msgId uint32, data []byte) error {
	// v0.3 实现
	//_, err := s.Conn.Write(data)
	//if err != nil {
	//	return err
	//}

	// v0.5 增加封包
	if s.isClosed {
		return errors.New("connection closed when send msg")
	}

	// step1: 生成封包对象
	dp := NewDataPack()
	// step2: 生成msg
	msg := Message{
		Id:      msgId,
		DataLen: len(data),
		Data:    data,
	}

	// step3: 将msg打包为字节流
	binaryData, err := dp.Pack(&msg)
	if err != nil {
		return err
	}

	// step4: 发送
	_, err = s.Conn.Write(binaryData)
	if err != nil {
		return err
	}

	return nil
}
