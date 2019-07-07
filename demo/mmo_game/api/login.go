package api

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
	"github.com/phuhao00/zinx/demo/mmo_game/proto/PB"
	"github.com/phuhao00/zinx/itface"
	"github.com/phuhao00/zinx/znet"
	"time"
	"github.com/phuhao00/zinx/demo/mmo_game/db"
	."github.com/phuhao00/zinx/demo/mmo_game/core"
)

type Login struct {
	znet.BaseRouter
}
const (
	USERNAME = "root"
	PASSWORD = "huhao123"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "rose_pd"
)
var MysqlDB *sql.DB
func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	MysqlDB,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return
	}
	MysqlDB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	MysqlDB.SetMaxOpenConns(100)//设置最大连接数
	MysqlDB.SetMaxIdleConns(16) //设置闲置连接数
}

func (*Login)SignUp(request itface.IRequest)  {
	//1. 将客户端传来的proto协议解码
	msg := &PB.Login{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Move: Position Unmarshal error ", err)
		return
	}
	strSelect:="select * from user where name="

	stmt,err:=MysqlDB.Prepare(strSelect)
	defer stmt.Close()
	if err==nil {
		rows,err:=stmt.Query(msg.Account)
		if err==nil {
			var userInfo db.UserTB
			if rows !=nil{
				rows.Scan(
					&userInfo.Id,
					&userInfo.Age,
					&userInfo.Email,
					&userInfo.Phone_num,
					&userInfo.Pwd,
					&userInfo.Sex,
					)
				player:=userInfo.ToPlayer()
				player.Conn=request.GetConnection()
				WorldMgrObj.AddPlayer(player)
				resp:=&PB.LoginResp{
					Code: 200,
					Desc: "login success",
				}
				player.SendMsg(request.GetMsgID(),resp)
			}
		}
	}
}
