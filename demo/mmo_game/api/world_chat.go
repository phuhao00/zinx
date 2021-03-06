package api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/phuhao00/zinx/demo/mmo_game/core"
	"github.com/phuhao00/zinx/demo/mmo_game/proto/PB"
	"github.com/phuhao00/zinx/itface"
	"github.com/phuhao00/zinx/znet"
	//"github.com/gin-gonic/gin"
)

//世界聊天 路由业务
type WorldChatApi struct {
	znet.BaseRouter
}
//
func (*WorldChatApi) Handle(request itface.IRequest) {
	//1. 将客户端传来的proto协议解码
	msg := &PB.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Talk Unmarshal error ", err)
		return
	}

	//2. 得知当前的消息是从哪个玩家传递来的,从连接属性pid中获取
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("GetProperty pid error", err)
		request.GetConnection().Stop()
		return
	}
	//3. 根据pid得到player对象
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	//4. 让player对象发起聊天广播请求
	player.Talk(msg.Content)
}