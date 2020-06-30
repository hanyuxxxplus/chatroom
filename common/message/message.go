package message

// 确定消息的值
const (
	LoginMesType = "loginMes"
	LoginResMesType = "loginResMes"
	RegisterMesType = "registerMes"
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
	Error string `json:"error"`// 错误信息
}


// 注册返回的消息
type RegisterMessage struct {

}