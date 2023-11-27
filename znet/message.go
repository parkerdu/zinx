package znet

type Message struct {
	Id      uint32 // 消息id
	DataLen int    // 消息长度
	Data    []byte // 有用的消息内容
}

// 获取消息属性
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() int {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}
func (m *Message) SetMsgLen(len int) {
	m.DataLen = len
}
func (m *Message) SetData(data []byte) {
	m.Data = data
}
