package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义一个userDao结构体
// 完成对user结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

var (
	MyUserDao *UserDao
)

func NewUserDao(pool *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{
		pool: pool,
	}

	return
}

// 根据用户id 返回一个User实例 + err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error){
	res , err := redis.String(conn.Do("HGet","users",id))
	if err != nil {
		if err == redis.ErrNil { // 表示在 users 哈希中，没有找到对应数据
			err = ERROR_USER_NOT_EXSIST
		}
		return
	}

	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json unmarshal err = " , err)
		return
	}

	return

}

func (this *UserDao) UserLogin(userId int,userPwd string)(user *User,err error){
	//获取一个redis pool 的连接
	conn := this.pool.Get()

	defer conn.Close()
	// 通过getUserById 获取user信息
	user , err = this.getUserById(conn,userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return


}

func (this *UserDao) UserRegister(mes *message.RegisterMessage)(err error){
	//获取一个redis pool 的连接
	conn := this.pool.Get()
	defer conn.Close()
	// 通过getUserById 获取user信息
	user , err := this.getUserById(conn,mes.User.UserId)
	if user != nil {
		err = ERROR_USER_EXSIST
		return
	}

	res , err  := json.Marshal(mes.User)

	_,err = conn.Do("HSet","users",mes.User.UserId,res)

	return
}