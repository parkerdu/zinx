package zinterface

/*
针对Message进行TLV的拆包
step1: 先读取header 得到消息长度和消息id
step2: 接着读取固定长度得到消息内容
*/

/*
封包、拆包、模块
面试tcp数据流，处于粘包问题
*/

type IDataPack interface {
	// 获取包的header的长度
	GetHeaderLen() int
	// 封包：把Message的结构封装成字节流
	Pack(msg IMessage) ([]byte, error)
	// 拆包：根据收到的字节流，将数据反序列化为Message格式
	UnPack(data []byte) (IMessage, error)
}
