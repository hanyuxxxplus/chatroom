package main

import (
	"chatroom/common/message"
	"chatroom/service/process"
	"chatroom/service/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 功能：更具客户端发送的消息种类不同，决定改用那个函数处理
func (this *Processor)serverProcessMes(mes *message.Message)(err error){
	switch mes.Type {
	case message.LoginMesType:// 处理登录的逻辑
		// 创建一个userProcessor实例
		up := &process.UserProcessor{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:// 处理注册的逻辑

	default:
		fmt.Println("消息类型不存在")
	}
	return
}

func (this *Processor) mesProcessor()(err error){

	// 读取客户端发送的信息
	ts := utils.Transfer{
		Conn: this.Conn,
	}
	for{
		mes,err := ts.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("服务器端正常关闭")
				return err
			}
			fmt.Println("for循环 readPkg返回的err 不等于 nil err =",err)
			return err
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes 函数处理返回error err = ",err)
			return err
		}
	}

}