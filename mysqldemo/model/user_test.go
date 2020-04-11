package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试开始: ")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试user相关的方法")
	//t.Run("测试添加用户：", testUser_AddUser)
	t.Run("测试查询用户:", testUser_GetUserById)
	//t.Run("测试查询用户:", testUser_GetUsers)
}

func testUser_AddUser(t *testing.T) {
	fmt.Println("test add user ====")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}

func testUser_GetUserById(t *testing.T) {
	fmt.Println("测试一条查询记录:")
	user := User{id: 55}
	u, _ := user.GetUserById()
	fmt.Println(*u)
}

func testUser_GetUsers(t *testing.T) {
	fmt.Println("测试所有查询记录:")
	user := User{}
	users, _ := user.FindUserInfo()
	fmt.Println(users[0])
	fmt.Println(&users[0])
	//遍历切片
	for k, v := range users {
		fmt.Println(k, v)
	}
}
