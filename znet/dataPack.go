package znet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"zinx/config"
	"zinx/zinterface"
)

const HeaderLen int = 8 // 前4个字节存放消息长度，后4个字节存放id

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的header的长度
func (d *DataPack) GetHeaderLen() int {
	return HeaderLen
}

// 封包：把Message的结构封装成字节流
func (d *DataPack) Pack(msg zinterface.IMessage) ([]byte, error) {
	// step1: 创建一个存放bytes字节流的缓冲
	buf := bytes.NewBuffer([]byte{})

	// step2: 写入header 中前4个字节，dataLen
	if err := binary.Write(buf, binary.BigEndian, uint32(msg.GetMsgLen())); err != nil {
		return nil, err
	}
	// step3: 写后四字节，id
	if err := binary.Write(buf, binary.BigEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// step4: 写史记数据
	if err := binary.Write(buf, binary.BigEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 拆包：输入为数据头header的字节，将数据反序列化为Message格式, 得到消息头，无法真实数据
func (d *DataPack) UnPack(data []byte) (zinterface.IMessage, error) {
	// step1: 创建一个从二进制数据的ioReader
	buf := bytes.NewReader(data)
	msg := &Message{}
	// step1: 读取header
	var length uint32
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	msg.DataLen = int(length)

	// step2: 从header中拿出dataLen + id
	if err := binary.Read(buf, binary.BigEndian, &msg.Id); err != nil {
		return nil, err
	}

	if config.MaxPackageSize() > 0 && msg.DataLen > config.MaxPackageSize() {
		return nil, fmt.Errorf("too large msg data len")

	}

	return msg, nil
}
