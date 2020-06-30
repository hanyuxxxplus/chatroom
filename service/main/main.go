package main

import (
	"fmt"
	"net"
)

// 处理和客户端的通讯
func doProcess(conn net.Conn){

	// 延时关闭conn
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("服务端延时关闭conn出错 err = ",err)
		}
	}()

	processor := &Processor{
		Conn: conn,
	}
	err := processor.mesProcessor()
	if err != nil {
		fmt.Println("客户端和服务器端通讯协程错误！", err)
		return
	}
}

func main(){
	fmt.Println("服务器8889")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil {
		fmt.Println("new.listen err = ",err)
		return
	}
	defer listen.Close()
	// 连接成功后，等待客户端连接服务器
	for  {
		fmt.Println("等待客户连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err = ",err)
		}
		// 一旦连接成功则启动一个协程，保持和客户端的通讯
		go doProcess(conn)
	}
}