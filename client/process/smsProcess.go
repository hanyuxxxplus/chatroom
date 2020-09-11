package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
)

type SmsProcess struct {

}

func (this *SmsProcess) sendGroupMes (content string)(err error) {
	var smsMes message.SmsMessage

	smsMes.Content = content
	smsMes.FromUserId = CUser.UserId

	data,err := json.Marshal(smsMes)

	if err != nil {
		return
	}


	var mes message.Message

	mes.Type = message.SmsMesType

	mes.Data = string(data)

	data , err  = json.Marshal(mes)

	if err != nil {
		return
	}

	ts := utils.Transfer{
		Conn: CUser.Conn,
	}

	err = ts.WritePkg(data)

	return err
}