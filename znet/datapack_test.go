package znet

import (
	"net"
	"testing"
)

func TestDataPack_Pack(t *testing.T) {
	/*
		模拟服务器
	*/
	// step1: 创建sockerTcp
	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		t.Fatal(err)
	}

	// step2: 处理客户端数据
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				t.Fatal(err)
			}

			// 处理链接
			go func() {
				// 拆包
				d := NewDataPack()
				for {
					// step1: 读取header
					header := make([]byte, HeaderLen)
					_, err := conn.Read(header)
					if err != nil {
						t.Fatal(err)
					}
					// byte 如何转化为 uint32? todo 是需要binary包来辅助的
					msg, err := d.UnPack(header)
					if err != nil {
						t.Fatal(err)
					}

					// step2: 通过header 拿真实数据
					if msg.GetMsgLen() > 0 {
						msg := msg.(*Message)
						// 开辟消息空间
						msg.Data = make([]byte, msg.GetMsgLen())

						offset := 0
						for offset < msg.GetMsgLen() {
							n, err := conn.Read(msg.Data)
							if err != nil {
								t.Fatal(err)
							}
							offset += n
						}

						// 打印消息查看
						t.Logf("len: %v, id: %d, data: %v", msg.GetMsgLen(), msg.GetMsgId(), string(msg.GetData()))
					}
				}

			}()
		}
	}()

	// step3: 客户端向上面服务端发送请求

	conn, err := net.Dial("tcp", ":7777")
	if err != nil {
		t.Fatal(err)
	}

	// 创建封包
	d := NewDataPack()

	// 模拟2个包粘一起

	msg1 := Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'}, // go 中单引号表示字符
	}
	data1, err := d.Pack(&msg1)
	if err != nil {
		t.Fatal(err)
	}

	msg2 := Message{
		Id:      1,
		DataLen: 2,
		Data:    []byte{'n', 'i'}, // go 中单引号表示字符
	}
	data2, err := d.Pack(&msg2)
	if err != nil {
		t.Fatal(err)
	}
	data1 = append(data1, data2...)

	conn.Write(data1)
	// 阻塞
	select {}

}
