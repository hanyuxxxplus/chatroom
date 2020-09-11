package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
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
	var content string
	switch key {
	case 1:
		showOnlineUser()
	case 2:
		fmt.Println("请输入要发送的消息")
		fmt.Scanln(&content)
		smsProcessor := SmsProcess{}
		err := smsProcessor.sendGroupMes(content)
		if err != nil {
			fmt.Println(err)
		}
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
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data),&notifyMes)
			if err != nil {
				fmt.Println(err)
				return
			}
			updateUserStatus(notifyMes.UserId,notifyMes.Status)
			showOnlineUser()
		case message.SmsMesType:
			var smsMes message.SmsMessage
			err = json.Unmarshal([]byte(mes.Data),&smsMes)
			if err != nil {
				fmt.Println(err)
				return
			}
			ShowSmsResMes(&smsMes)
		default:
			fmt.Println("无法识别的类型")
		}

	}
}
