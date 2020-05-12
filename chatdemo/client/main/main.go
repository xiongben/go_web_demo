package main

import (
	"fmt"
	clientprocess "go_web_demo/chatdemo/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	for true {
		fmt.Println("=======欢迎登陆多人聊天系统==========")
		fmt.Println("1 login the chatroom")
		fmt.Println("2 register the chatroom")
		fmt.Println("3 out the chatroom")
		fmt.Println(" please choose 1 - 3")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("login in chatroom")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%v\n", &userPwd)
			up := &clientprocess.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("register")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入设置密码")
			fmt.Scanf("%v\n", &userPwd)

			fmt.Println("请输入用户名称")
			fmt.Scanf("%v\n", &userName)

			up := &clientprocess.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("login out")
			//loop = false
		default:
			fmt.Println("you put error message")

		}

	}
}
