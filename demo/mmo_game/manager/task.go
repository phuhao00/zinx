package manager

import (
	"encoding/json"
	. "github.com/phuhao00/zinx/demo/mmo_game/conf"
	"io/ioutil"
)

//
type  TaskManager struct {
	tasks map[int32]Task //
}
//
func NewTaskManager()*TaskManager{
	TaskMgr:=&TaskManager{
		tasks:make(map[int32]Task),
	}
return TaskMgr
}
//
func (TaskM *TaskManager)LoadConfigs() {
	f, err := ioutil.ReadFile(TaskConfigPath)
	if err != nil {
		panic(err)
	}
	var rows []*TaskConfig
	err = json.Unmarshal(f, &rows)
	if err != nil {
		panic(err)
	}
	TaskM.tasks = make(map[int32]Task)
	for _, row := range rows {
		if row.Classification==1 {
			tmp := &TaskNormal{
				State:          TaskState_Init,
				Classification: row.Classification,
				IsDone:         false,
			}
			var tmpTargets []*Target
			for _, ele := range row.Data {
				tmpTarget := &Target{
					Type:   ele[0],
					Data:   ele,
					Desc:   row.Desc,
					IsDone: false,
				}
				GlobalMgr.TaskEM.AddSubscriber(tmpTarget)
				tmpTargets = append(tmpTargets, tmpTarget)
			}
			tmp.Targets=tmpTargets
			TaskM.tasks[row.ID] =tmp
		}
	}
}

func (TaskM *TaskManager)Open()  {

}
func (TaskM *TaskManager)Close()  {

}
func (TaskM *TaskManager)IsClose()  {

}
func (TaskM *TaskManager)Before()  {
	TaskM.LoadConfigs()
}
func (TaskM *TaskManager)After()  {

}
