package conf

import (
	. "github.com/gookit/event"
	. "github.com/phuhao00/zinx/demo/mmo_game/pb/SubscriberID"
)

var subscribedEvents =map[SubscriberId]map[string]interface{}{
	SubscriberId_Normal:{ SubscriberId_Normal.String(): HasMoney.Handle},
}

type HasMoney struct {

}
//
func (HasMoney)Handle(e Event) error  {
	return  nil
}



