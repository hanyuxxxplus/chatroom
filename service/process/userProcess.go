package process

import (
	"encoding/json"
	"fmt"
	"net"
	"project01/chat/common/message"
	"project01/chat/service/utils"
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

	// 声明一个loginMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "abc"{
		loginResMes.Code = 200
	}else{
		loginResMes.Code = 500
		loginResMes.Error = "用户名不合法"
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