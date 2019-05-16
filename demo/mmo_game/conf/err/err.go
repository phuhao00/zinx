package err

import (
	"strconv"
)
//
type ErrorInfo struct {
	Code uint64
	Desc string
}
//
func (self *ErrorInfo)String() string  {
 	str:="Code:"+strconv.FormatUint(self.Code,64)+"Desc:"+self.Desc
	return str
}
var(
	ErrRequest=&ErrorInfo{ERR_RequestParam,""}
	ErrMoneyNotEnough=&ErrorInfo{ERR_MoneyNotEnough,""}
)




