package manager

import ."github.com/gookit/event"

//
type GlobalManager struct {
	TaskEM *Manager
	Managers []managerInterface
}
var GlobalMgr=&GlobalManager{}
//
func (self *GlobalManager)Open()  {
	self.TaskEM=NewManager("Task")
}
//
func (self *GlobalManager)Close()  {

}
//
func (self *GlobalManager)IsClose()  {

}
//
func (self *GlobalManager)Before()  {
	self.Managers=[]managerInterface{
		NewTaskManager(),
	}
	for i:=0;i< len(self.Managers);i++{
		self.Managers[i].Before()
	}
}
//
func (self *GlobalManager)After()  {

}




