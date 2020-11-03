package main
import (
	"chatroom/client/process"
	"fmt"
)

// 定义两个全局标量 一个用户id 一个用户密码
var userId int
var userPwd string
var userName string
func main(){
	var key int
	var loop bool= true
	for{
		fmt.Println("---------------1欢迎登录多人聊天系统---------------")
		fmt.Printf("\t\t\t 1、登录聊天室\n")
		fmt.Printf("\t\t\t 2、注册用户\n")
		fmt.Printf("\t\t\t 3、推出系统\n")
		fmt.Printf("\t\t\t 请选择（1-3）：\n")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登录聊天系统")
			fmt.Println("请输入用户ID")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)
			// 完成登录
			up := process.UserProcessor{}
			err := up.Login(userId,userPwd)
			if err != nil {
				fmt.Println("main:",err)
			}
		case 2:
			fmt.Println("请输入用户ID")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户昵称")
			fmt.Scanln(&userName)
			// 完成登录
			up := process.UserProcessor{}
			err := up.Register(userId,userPwd,userName)
			if err != nil {
				fmt.Println("main:",err)
			}else{
				fmt.Println("注册成功，请重新登录")
			}
		case 3:
			fmt.Println("推出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}

		if !loop {
			break
		}
	}
}