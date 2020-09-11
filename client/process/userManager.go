package process

import (
	clientModel "chatroom/client/model"
	"chatroom/common/model"
	"fmt"
)

var userManager = make(map[int]*model.User , 1024)
var CUser clientModel.CurrentUser

func updateUserStatus(userId int,status int) {
	user,ok := userManager[userId]
	if !ok {
		user = &model.User{UserId: userId}
	}

	user.UserStatus = status
	userManager[userId] = user
}

func showOnlineUser(){
	for k , _ := range userManager {
		fmt.Println("当前在线用户： ",k)
	}
}