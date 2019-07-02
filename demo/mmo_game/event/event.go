package main

import (
	"github.com/gookit/event"
	. "github.com/phuhao00/zinx/demo/mmo_game/conf"
)
func main()  {
	//user:=&User{
	//
	//}
	user:=NewUser()

	e:=event.NewBasic("task1",nil)
	e.SetTarget("2")
	EM_task.AddEvent(e)
	var listener event.ListenerFunc= func(e event.Event) error {
		if e.(*event.BasicEvent).Target().(string)=="2" {
			e.Abort(true)
		}
		return nil
	}

	EM_task.AddListener("",listener,event.AboveNormal)

	EM_task.FireEvent(e)
}

//初始化任务
func NewUser() *User {
	user:=&User{
		Tasks:nil,
	}
	return user
}

