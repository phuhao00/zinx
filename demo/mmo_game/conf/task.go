package conf

import ."awesomeProject/demo/mmo_game/pb/SubscriberID"
//任务的状态
const (
	TaskState_Init =iota+1
	TaskState_Done
)
//任务配置的路径
const TaskConfigPath ="configs/task.json"
//
type TaskConfig struct {
	ID int32
	Classification int32
	Data [][]int32
	Desc string
}
//
type Target struct {
	Type int32
	Data []int32
	Desc string
	IsDone bool
}
//
func (task *Target)SubscribedEvents() map[string]interface{} {
	return subscribedEvents[SubscriberId(task.Type)]
}
//
type Task interface {
	SetDone()
}
//
type TaskNormal struct {
	Id	int32 //任务ID
	State int32 //任务状态
	Classification int32 //任务类别，用于显示等
	IsDone bool
	Targets []*Target
}
//
func (task *TaskNormal) SetDone(){
	task.IsDone=true
	task.State=TaskState_Done
}


