package message

import "chatroom/common/model"
// 确定消息的值
const (
	LoginMesType = "loginMes"
	LoginResMesType = "loginResMes"
	RegisterMesType = "registerMes"
	RegisterResMessageType = "registerResMessage"
	NotifyUserStatusMesType = "notifyUserStatusMes"
	SmsMesType = "smsMesType"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 登录消息
type LoginMessage struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

// 登录返回的消息
type LoginResMessage struct {
	Code int `json:"code"`// 状态码 500：表示该用户未注册  200：表示登录成功
	OnlineUserList []int `json:"onlineUserList"` // 在线用户列表
	Error string `json:"error"`// 错误信息
}


// 注册返回的消息
type RegisterMessage struct {
	User model.User `json:"user"`
}

type RegisterResMessage struct {
	Code int `json:"code"`// 状态码 400：表示该用户已注册  200：表示注册成功
	Error string `json:"error"`// 错误信息
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMessage struct {
	Content string `json:"content"`
	FromUserId int `json:"fromUserId"`
}