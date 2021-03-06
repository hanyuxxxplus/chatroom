package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"chatroom/common/model"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcessor struct {

}

func (this *UserProcessor)Register(userId int, userPwd string, userName string) (err error) {
	// 开始定制协议
	// 1、连接到服务器
	conn ,err := net.Dial("tcp","localhost:8889") // localhost:8889 应该做成一个配置文件
	if err!= nil {
		return
	}
	defer func() {
		closeErr := conn.Close()
		if closeErr != nil {
			err = closeErr
			return
		}
	}()
	// 2、通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	// 3 、创建一个registerMessage 结构体
	var user = model.User{
		UserId: userId,
		UserPwd: userPwd,
		UserName: userName,
	}
	registerMes := message.RegisterMessage{
		User: user,
	}
	// 4、序列化loginMessage
	data , err := json.Marshal(registerMes)
	if err != nil {
		return
	}

	// 5、将数据赋值给data
	mes.Data = string(data)
	fmt.Println(mes.Data)
	// 6、将mes 序列化
	data , err = json.Marshal(mes)
	if err != nil {
		return
	}

	// 发送数据
	ts := &utils.Transfer{
		Conn: conn,
	}

	err = ts.WritePkg(data)
	if err != nil {
		return
	}
	// 接受返回的code
	resMes , err := ts.ReadPkg()
	if err != nil {
		return
	}
	// 将resMes data部分反序列化 LoginResMes
	var registerResMes message.RegisterResMessage
	err = json.Unmarshal([]byte(resMes.Data),&registerResMes)
	if err != nil {
		return
	}
	if registerResMes.Code == 200{
		// 在客户端启动一个协程
		// 该协程用来保持和服务器端的通讯
		// 如果服务器端有消息推送给客户端，则接收并显示
		go serverProcessMes(conn)
		for{
			ShowMenu()
		}
	}else{
		err = errors.New(registerResMes.Error)
		return
	}
}

func (this *UserProcessor)Login(userId int, userPwd string) (err error) {
	// 开始定制协议
	// 1、连接到服务器
	conn ,err := net.Dial("tcp","localhost:8889") // localhost:8889 应该做成一个配置文件
	if err!= nil {
		return
	}
	defer func() {
		closeErr := conn.Close()
		if closeErr != nil {
			err = closeErr
			return
		}
	}()
	// 2、通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3 、创建一个loginMessage 结构体
	loginMes := message.LoginMessage{
		UserId: userId,
		UserPwd: userPwd,
		UserName: "tom",
	}
	// 4、序列化loginMessage
	data , err := json.Marshal(loginMes)
	if err != nil {
		return
	}

	// 5、将数据赋值给data
	mes.Data = string(data)
	fmt.Println(mes.Data)
	// 6、将mes 序列化
	data , err = json.Marshal(mes)
	if err != nil {
		return
	}

	// 7、此时开始准备发送数据
	// 7.1 先将长度发送给服务器
	// 先获取data的长度 =》 转成一个表示长度的切片
	//var pkgLen uint32
	//pkgLen = uint32(len(data))
	//
	//dataBytes := make([]byte,4)
	//binary.BigEndian.PutUint32(dataBytes,pkgLen)
	//n , err := conn.Write(dataBytes)
	//if n != 4 || err != nil {
	//	fmt.Println("conn.Write() err = ",err)
	//	return
	//}
	//fmt.Println("客户端发送消息成功 消息长度=",len(data))
	//
	//// 发送消息本身
	//_ , err = conn.Write(data)
	//if err != nil {
	//	fmt.Println("conn.Write() data err = ",err)
	//	return
	//}
	// 发送数据
	ts := &utils.Transfer{
		Conn: conn,
	}

	err = ts.WritePkg(data)
	if err != nil {
		return
	}
	// 接受返回的code
	resMes , err := ts.ReadPkg()
	if err != nil {
		return
	}
	// 将resMes data部分反序列化 LoginResMes
	var loginResMes message.LoginResMessage
	err = json.Unmarshal([]byte(resMes.Data),&loginResMes)
	if err != nil {
		return
	}

	if loginResMes.Code == 200{
		// 在客户端启动一个协程
		// 该协程用来保持和服务器端的通讯
		// 如果服务器端有消息推送给客户端，则接收并显示
		for _,v := range loginResMes.OnlineUserList {
			if v == userId {
				continue
			}
			updateUserStatus(v,model.UserOnline)
		}
		CUser.Conn = conn
		CUser.UserId = userId
		CUser.UserStatus = model.UserOnline
		showOnlineUser()
		go serverProcessMes(conn)
		for{
			ShowMenu()
		}
	}else if loginResMes.Code == 500 {
		err = errors.New(loginResMes.Error)
		return

	}
	return
}