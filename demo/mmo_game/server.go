package main

import (
	"flag"
	"fmt"
	"github.com/phuhao00/zinx/demo/mmo_game/api"
	"github.com/phuhao00/zinx/demo/mmo_game/core"
	"github.com/phuhao00/zinx/itface"
	"github.com/phuhao00/zinx/znet"
	"github.com/phuhao00/zinx/demo/mmo_game/proto/ID"
)

//当客户端建立连接的时候的hook函数
func OnConnecionAdd(conn itface.IConnection)  {
	//创建一个玩家
	player := core.NewPlayer(conn)

	//同步当前的PlayerID给客户端， 走MsgID:1 消息
	player.SyncPid()

	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	player.BroadCastStartPosition()

	//将当前新上线玩家添加到worldManager中
	core.WorldMgrObj.AddPlayer(player)

	//将该连接绑定属性Pid
	conn.SetProperty("pid", player.Pid)

	//同步周边玩家上线信息，与现实周边玩家信息
	player.SyncSurrounding()

	fmt.Println("=====> Player pidId = ", player.Pid, " arrived ====")
}

//当客户端断开连接的时候的hook函数
func OnConnectionLost(conn itface.IConnection) {
	//获取当前连接的Pid属性
	pid, _ := conn.GetProperty("pid")

	//根据pid获取对应的玩家对象
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	//触发玩家下线业务
	if pid != nil {
		player.LostConnection()
	}

	fmt.Println("====> Player ", pid , " left =====")

}

var isDaemon, isForever, isOnlyCheckAlive bool
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
func init()  {
	 flag.BoolVar(&isDaemon,"daemon", false, "is daemon background")
	 flag.BoolVar(&isForever,"forever", false, "is guard forever")
	 flag.BoolVar(&isOnlyCheckAlive,"checkAlive", false, "is guard forever")
}
func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//注册客户端连接建立和丢失函数
	s.SetOnConnStart(OnConnecionAdd)
	s.SetOnConnStop(OnConnectionLost)

	//注册路由
	s.AddRouter(ID.Message_Chat, &api.WorldChatApi{})
	s.AddRouter(ID.Message_Move, &api.MoveApi{})
	s.AddRouter(ID.Message_SignIn, &api.Login{})
	//启动服务
	s.Serve()
}
