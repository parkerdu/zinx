package znet

import (
	"fmt"
	"go/token"
	"net"
	"zinx/zinterface"
)

// 定义一个实现了server接口的类
type Server struct {
	// 服务名称
	Name string
	// 服务绑定的ip版本，例如tcp
	IPVersion string
	// ip
	IP string
	// 端口
	Port int
}

func (s *Server) Start() {
	fmt.Printf("server listen %s:%v", s.IP, s.Port)

	// 1、获取一个tcp的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println(err)
		return
	}


	// 2、监听服务器地址
	go func() {
		listen, err  := net.ListenTCP(s.IPVersion, addr)
		// 3、阻塞等待客户端链接，并根据路由处理业务
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				return
			}


			// 读取链接中内容
			session := NewConnection(conn, )


		}

	}()


}

func (s *Server) Stop() {

}

// 用户层只调用该方法
func (s *Server) Server() {

}

// 初始化server
func NewServer(name string) zinterface.IServer {
	s := Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8089,
	}
	return &s
}

func dealCollback(conn *net.TCPConn, data []byte, n int) error {
	if _, err := conn.Write(data[:n]); err != nil {
		fmt.Println(err)
		continue
	}
	return nil
}



func ()  {

}