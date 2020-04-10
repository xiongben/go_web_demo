package model

import (
	"fmt"
	"testing"
)

func TestUser_AddUser(t *testing.T) {
	fmt.Println("test add user ====")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}
