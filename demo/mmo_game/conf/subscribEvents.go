package conf

import (
	. "awesomeProject/demo/mmo_game/pb/SubscriberID"
	."github.com/gookit/event"
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



