package clientmodel

import (
	"go_web_demo/chatdemo/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
