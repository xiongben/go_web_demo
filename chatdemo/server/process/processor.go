package process

import (
	"fmt"
	"go_web_demo/chatdemo/common/message"
	"go_web_demo/chatdemo/server/utils"
	"io"
	"net"
)

type ProcessorObj struct {
	Conn net.Conn
}

func (this *ProcessorObj) SeverProcessMess(mes *message.Message) (err error) {
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		up := &UserProcess{
			Conn: this.Conn,
		}
		_, err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("message type is not exit,can't deal it!")
	}
	return
}

func (this *ProcessorObj) SeverProcess() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
			Buf:  [8096]byte{},
		}
		mess, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("server is out and the client is out too..")
				return err
			} else {
				fmt.Println("readpkg err = ", err)
				return err
			}
		}
		err = this.SeverProcessMess(&mess)
		if err != nil {
			return err
		}
	}
}
