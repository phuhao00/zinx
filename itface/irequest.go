package itface

import PB "github.com/phuhao00/zinx/demo/mmo_game/proto/ID"

/*
	IRequest 接口：
	实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface{
	GetConnection() IConnection	//获取请求连接信息
	GetData() []byte			//获取请求消息的数据
	GetMsgID() PB.Message			//获取请求的消息ID
}

