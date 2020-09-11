package process

import (
	"chatroom/common/message"
	"chatroom/service/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendMessageToOnlineUser(mes *message.Message){
	// 反序列化 客户端传过来的数据
	var smsMes message.SmsMessage
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("SendMessageToOnlineUser unmarshal err =",err)
		return
	}
	data , err := json.Marshal(mes)

	if err != nil {
		fmt.Println("Marshal err = ",err)
		return
	}
	// 发送给在线的学生
	for k,v := range Umg.userList {
		if k == smsMes.FromUserId {
			continue
		}

		this.SendMessageToSingleUser(data,v.Conn)
	}

}

func (this *SmsProcess) SendMessageToSingleUser(data []byte,conn net.Conn){
	var ts = &utils.Transfer{
		Conn: conn,
	}

	err := ts.WritePkg(data)

	if err != nil {
		fmt.Println("SendMessageToSingleUser WritePkg err = ",err)
	}
	return
}