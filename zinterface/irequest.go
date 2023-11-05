package zinterface

// 把封装的Iconnection 和 请求包的数据封装在一起

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte
}
