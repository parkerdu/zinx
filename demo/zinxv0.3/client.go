package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/

func main() {
	fmt.Println("client start...")
	// 1、链接服务，得到conn
	conn, err := net.Dial("tcp", "127.0.0.1:8079")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2、往conn中写入数据
	for {
		_, err := conn.Write([]byte("Hello Zinx v0.1"))
		if err != nil {
			fmt.Println("write conn err: ", err)
			return
		}

		// 读取conn中的response
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("收到response: ", string(buf[:n]))

		time.Sleep(time.Second)
	}
}
