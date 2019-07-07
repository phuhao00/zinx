package db

import (
	."github.com/phuhao00/zinx/demo/mmo_game/core"
)

type UserTB struct {
	Id int64
	Name string
	Pwd string
	Sex int
	Age int
	Phone_num  string
	Email string
}

func (self *UserTB)ToPlayer()  *Player {
	player:=&Player{
		Pid:int32( self.Id ),
	}
	return  player
}
func (self *UserTB)Run()  {


}