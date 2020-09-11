package process

import (
	"chatroom/common/message"
	"chatroom/service/model"
	"chatroom/service/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn net.Conn
}

// 登录请求处理函数
func (this *UserProcessor)ServerProcessLogin(mes *message.Message)(err error){
	var loginMes message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err ", err)
	}
	// 先声明一个 responseResMessage
	var res message.Message
	res.Type = message.LoginResMesType
	// 声明一个loginResMes
	var loginResMes message.LoginResMessage

	user, err := model.MyUserDao.UserLogin(loginMes.UserId,loginMes.UserPwd)
	// 声明一个loginMes
	if err != nil {
		loginResMes.Code = 500
		loginResMes.Error = err.Error()
	}else{
		loginResMes.Code = 200
		Umg.AddUserOnline(user.UserId,this)
		// 将在线用户返回给登录用户
		for k , _ := range Umg.userList {
			loginResMes.OnlineUserList = append(loginResMes.OnlineUserList,k)
		}
		// TODO 通知其他用户该用户上线事件，应该可以放到协程中去完成
		this.NotifyOtherUsers(user.UserId , model.UserOnline)
	}
	var loginResMesByte []byte
	loginResMesByte ,err = json.Marshal(loginResMes)
	if err != nil {
		return
	}

	res.Data = string(loginResMesByte)

	sendData ,err := json.Marshal(res)
	if err != nil {
		return
	}

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}

	return transfer.WritePkg(sendData)
}

func (this *UserProcessor) NotifyOtherUsers(userId int , status int){
	var onlineUserList = Umg.GetAllOnlineUser()
	for k , _ :=range onlineUserList{
		if k == userId {
			continue
		}

		this.NotifyUserStatus(k , status, userId)
	}
}

func (this *UserProcessor) NotifyUserStatus(userId int, status int, onlineUserId int) {
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = onlineUserId
	notifyUserStatusMes.Status = status

	 data , err  := json.Marshal(notifyUserStatusMes)

	 if err != nil {
		 return
	 }

	var mes message.Message

	 mes.Data = string(data)
	 mes.Type = message.NotifyUserStatusMesType

	 data ,err = json.Marshal(mes)
	 if err != nil {
		 return
	 }
	transfer := &utils.Transfer{
		Conn: Umg.userList[userId].Conn,
	}

	err = transfer.WritePkg(data)

	if err != nil {
		fmt.Println("transfer.WritePkg err = ",err)
	}
}

func (this *UserProcessor) ServerProcessRegister (mes *message.Message)(err error) {
	var registerMes message.RegisterMessage
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err ", err)
	}

	// 先声明一个 responseResMessage
	var res message.Message
	res.Type = message.RegisterResMessageType
	// 声明一个registerResMes
	var registerResMes message.RegisterResMessage
	err = model.MyUserDao.UserRegister(&registerMes)

	if err != nil {
		registerResMes.Code = 500
		registerResMes.Error = err.Error()
	}else{
		registerResMes.Code = 200
	}
	var registerResMesByte []byte
	registerResMesByte ,err = json.Marshal(registerResMes)
	if err != nil {
		return
	}

	res.Data = string(registerResMesByte)

	sendData ,err := json.Marshal(res)
	if err != nil {
		return
	}
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}

	return transfer.WritePkg(sendData)

}