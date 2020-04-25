package process

import (
	"go_web_demo/chatdemo/common/message"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

}
