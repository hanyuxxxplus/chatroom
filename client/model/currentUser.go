package model

import (
	"chatroom/common/model"
	"net"
)

type CurrentUser struct {
	Conn net.Conn
	model.User
}