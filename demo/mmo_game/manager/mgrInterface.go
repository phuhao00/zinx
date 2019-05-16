package manager

type managerInterface interface {
	Before()//
	Open()
	After()
	Close()
	IsClose()
}
