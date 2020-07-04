package model

import (
	"fmt",
	""
)

// 定义一个userDao结构体
// 完成对user结构体的各种操作
type UserDao struct {
	pool *redis.pool
}