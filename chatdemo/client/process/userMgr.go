package clientprocess

import (
	"fmt"
	clientmodel "go_web_demo/chatdemo/client/model"
	"go_web_demo/chatdemo/common/message"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser clientmodel.CurUser

func outPutOnlineUsers() {
	fmt.Println("online user list:")
	for id, _ := range onlineUsers {
		fmt.Println("user id : \t", id)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outPutOnlineUsers()
}
