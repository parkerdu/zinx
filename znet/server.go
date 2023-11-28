package znet

import (
	"fmt"
	"net"
	"zinx/config"
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
	// v0.3 增加路由字段, 该版本一个服务只能注册一个路由，暂不支持多路由
	//Route zinterface.IRouter

	// v0.6 msgHandle增加多个路由, 当前server的消息管理模块，绑定msgId和处理方法
	Handle zinterface.MsgHandler
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
			session := NewSession(conn, seqId, s.Handle)
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

// todo 可以直接改成 添加func(request zinterface.IRequest) []byte, 而不是每个理由都搞一个结构体
func (s *Server) AddRouter(msgId uint32, route zinterface.IRouter) {
	s.Handle.AddRoute(msgId, route)
}

// 初始化server
func NewServer(name string) zinterface.IServer {
	s := Server{
		Name:      config.Server().Name,
		IPVersion: "tcp4",
		IP:        config.Server().Host,
		Port:      config.Server().Port,
		Handle:    NewMsgHandle(),
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
