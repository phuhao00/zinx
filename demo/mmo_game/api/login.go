package api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/phuhao00/zinx/demo/mmo_game/proto/PB"
	"github.com/phuhao00/zinx/itface"
	"github.com/phuhao00/zinx/znet"
	"database/sql"
 	"github.com/go-sql-driver/mysql"

)

type sign_up struct {
	znet.BaseRouter
}

func (*sign_up)SignUp(request itface.IRequest)  {
	//1. 将客户端传来的proto协议解码
	msg := &PB.Login{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Move: Position Unmarshal error ", err)
		return
	}

	
}
