package znet

import (
	"fmt"
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
	fmt.Printf("server listen %s:%v\n", s.IP, s.Port)

	// 1、获取一个tcp的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	var seqId uint32 = 0

	// 2、监听服务器地址
	listen, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("start Zinx server named %s success, listening...", s.Name)
	// 3、阻塞等待客户端链接，并根据路由处理业务
	go func() {
		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				return
			}
			// 读取链接中内容, 并进行业务处理
			session := NewConnection(conn, seqId, dealCollback)
			go session.Start()
			seqId++
		}
	}()
}

func (s *Server) Stop() {
	// todo 回收资源，关闭链接等操作
	fmt.Printf("server %s stop", s.Name)
}

// 用户层只调用该方法来启动服务
func (s *Server) Server() {
	s.Start()

	// TODO 做一些启动服务之后的额外业务，而不是在start里面阻塞，这就是为啥还要用Server来包装
	// 阻塞等待退出服务
	select {}
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
		return err
	}
	return nil
}
