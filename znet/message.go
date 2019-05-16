package znet

import (
	"awesomeProject/demo/mmo_game/conf"
)

type Message struct {
	DataLen   uint32         //消息的长度
	Id        uint32         //消息的ID
	Data      []byte         //消息的内容
	Version   uint16         //暂时尚未使用。首位征用，用于表示是否有ReqId
	ErrorInfo conf.ErrorInfo //错误信息
}

//创建一个Message消息包
func NewMsgPackage(id uint32, data []byte,ErrorInfo *conf.ErrorInfo) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		Id:     id,
		Data:   data,
		ErrorInfo:*ErrorInfo,
	}
}

//获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//获取消息ID
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

//获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

//设计消息ID
func (msg *Message) SetMsgId(msgId uint32) {
	msg.Id = msgId
}

//设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
////
//func (msg *Message)Pack()  {
//
//}