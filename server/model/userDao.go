package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//操作redis数据库

// 定义UserDao结构体，完成对结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1.根据用户Id返回用户实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过id去数据库 查询
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err == redis.ErrNil {
		//没有查找到对应id
		err = ERROR_USER_NOTEXITS
		return
	}

	//找到对应信息，反序列化成User结构体
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("反序列化成User结构体失败，错误是：", err)
		return
	}
	return
}

//完成登录校验
//如果信息正确则返回user对象，否则返回err信息

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//从连接池获取一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//没有错误成功拿到user信息，进行对比
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
