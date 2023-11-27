package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

/*
模拟客户端
*/

func main() {
	fmt.Println("client start...")
	// 1、链接服务，得到conn
	conn, err := net.Dial("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2、往conn中写入数据
	for {
		// 将消息封装成msg包
		dp := znet.NewDataPack()
		data := []byte("Hello Zinx v0.5")
		msg := znet.Message{
			Id:      1,
			DataLen: len(data),
			Data:    data,
		}
		binaryData, err := dp.Pack(&msg)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		_, err = conn.Write(binaryData)
		if err != nil {
			fmt.Println("write conn err: ", err)
			return
		}

		// 读取conn中的response, 并拆包

		header := make([]byte, dp.GetHeaderLen())
		offset := 0
		for offset < dp.GetHeaderLen() {
			n, err := conn.Read(header[offset:])
			if err != nil {
				fmt.Println(err)
			}
			offset += n
		}

		msg1, err := dp.UnPack(header)
		if err != nil {
			fmt.Println(err)
		}

		buf := make([]byte, msg1.GetMsgLen())
		if _, err = io.ReadFull(conn, buf); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("收到response: ", string(buf[:msg1.GetMsgLen()]))
		time.Sleep(time.Second)
	}
}
