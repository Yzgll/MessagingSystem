package main

import (
	"MessageSystem/client/process"
	"fmt"
)

// 全局变量一个账号一个密码
var userId int
var userPwd string
var userName string

func main() {

	var key int

	for true {
		fmt.Println("************************************************************")
		fmt.Println("*                  欢迎登录多人聊天系统                    *")
		fmt.Println("************************************************************")
		fmt.Println("*                    1 登录系统                            *")
		fmt.Println("*                    2 注册用户                            *")
		fmt.Println("*                    3 退出系统                            *")
		fmt.Println("*                    请选择（1-3）                         *")
		fmt.Println("************************************************************")

		//输入操作
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("\t\t\t登录系统")
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			//loop = false
			//完成登录
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("\t\t\t注册用户")
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户昵称")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("\t\t\t退出系统")
		default:
			fmt.Println("输入有误请重新输入!")
		}
	}
	//处理用户操作
	// if key == 1 {
	// 	fmt.Println("请输入用户ID")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("请输入用户密码")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	login(userId, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败")
	// 	// } else {
	// 	// 	fmt.Println("登录成功")
	// 	// }
	// } else if key == 2 {
	// 	fmt.Println("用户注册")
	// }
}
