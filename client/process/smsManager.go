package process

import (
	"chatroom/common/message"
	"fmt"
)

type SmsManager struct {

}

func ShowSmsResMes (smsMes *message.SmsMessage) {
	content := fmt.Sprintf("用户ID（%d）对大家说：%s",smsMes.FromUserId,smsMes.Content)
	fmt.Println(content)
}