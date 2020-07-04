package model

import (
	"errors"
)

// 根据业务逻辑自定义错误
var (
	ERROR_USER_NOT_EXSIST = errors.New("该用户用户不存在")
	ERROR_USER_EXSIST = errors.New("用户已存在")
	ERROR_USER_PWD = errors.New("密码错误")
)