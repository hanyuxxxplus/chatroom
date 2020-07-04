package process

import (
	"chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

func ShowMenu(){
	fmt.Println("------你好，xxxx------")
	fmt.Println("------1、用户列表------")
	fmt.Println("------2、发送消息------")
	fmt.Println("------3、信息列表------")
	fmt.Println("------4、推出系统------")
	fmt.Println("请选择 1~4：")
	var key int
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("显示用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	}

}


func serverProcessMes(conn net.Conn){
	// 创建一个transfer实例，不停的读取服务器推送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}

	for  {
		fmt.Printf("客户端正在等待读取服务器端的消息")
		mes , err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.readPkg err = ",err)
			return
		}
		// 读取到消息则是下一步的处理逻辑

		fmt.Println(mes)

	}
}
