package process

import (
	"encoding/json"
	"fmt"
	"go_web_demo/chatdemo/common/message"
	"go_web_demo/chatdemo/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmashal err = ", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes
	//数据库操作
	loginResMes.Code = 200
	fmt.Println("user login success!")
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("jsonmarshal err = ", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("jsonmarshal err = ", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
