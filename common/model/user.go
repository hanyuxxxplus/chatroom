package model

// 定义一个用户的结构体
type User struct {
	// 为了保证反序列化的成功
	//  必须保证用户信息的json字符串的key 必须和结构体的tag名字一致
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}