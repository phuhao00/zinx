package main

import (
	"fmt"
	"net"
)

func main()  {
	addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", "0.0.0.0",8999))
	if err != nil {
		fmt.Println("resolve tcp addr err: ", err)
		return
	}
	conn, err:= net.DialTCP("tcp4",nil, addr)
	if err != nil {
		return
	}
	sender(conn)
}


func sender(conn net.Conn) {
	words := "Hello Server!"
	conn.Write([]byte(words))
	fmt.Println("send over")

	//接收服务端反馈
	buffer := make([]byte, 2048)
	_, err := conn.Read(buffer)
	if err != nil {
		return
	}
}