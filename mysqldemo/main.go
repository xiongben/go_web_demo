package main

import "go_web_demo/mysqldemo/model"

func main() {
	user := &model.User{}
	user.FindUserInfo()

}
