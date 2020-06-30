package process

import (
	"fmt"
	"os"
)

func ShowMenu(){
	fmt.Println("------你好，xxxx------")
	fmt.Println("------1、用户列表------")
	fmt.Println("------2、发送消息------")
	fmt.Println("------3、信息列表------")
	fmt.Println("------4、推出系统------")
	fmt.Println("请选择 1~4：")
	var key int
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("显示用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("推出系统")
		os.Exit(0)
	}

}