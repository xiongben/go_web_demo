package clientprocess

import (
	"fmt"
	"go_web_demo/chatdemo/common/message"
	"go_web_demo/chatdemo/server/utils"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("-----------user login success--------")
	fmt.Println("-----------1, show online user list----------")
	fmt.Println("-----------2, send message----------")
	fmt.Println("-----------3, message list----------")
	fmt.Println("-----------4, out system----------")
	fmt.Println("-----------choose (1-4)----------")
	var key int
	var content string

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("show user list")
	case 2:
		fmt.Println("say someting")
	case 3:
		fmt.Println("message list")
	case 4:
		fmt.Println("out system")
		os.Exit(0)
	default:
		fmt.Println("you put error choice")
	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	for {
		fmt.Println("client is reading message from server")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.readpkg err = ", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			fmt.Println("====")
		case message.SmsMesType:
			fmt.Println("=====")
		default:
			fmt.Println("server return unknow message type")
		}
	}
}
