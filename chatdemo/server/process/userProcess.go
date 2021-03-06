package process

import (
	"encoding/json"
	"fmt"
	"go_web_demo/chatdemo/common/message"
	"go_web_demo/chatdemo/server/model"
	"go_web_demo/chatdemo/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json unmashal err = ", err)
		return
	}
	fmt.Println("===regismes===")
	fmt.Println(registerMes)
	fmt.Println("===regismes===")
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	//数据库操作
	err = model.MyUserDao.Register(registerMes.User.UserId, registerMes.User.UserName, registerMes.User.UserPwd)
	fmt.Println("register sql err = ", err)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "server error ..."
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("user register success!")
	}

	data, err := json.Marshal(registerResMes)
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

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (user model.UserSql, err error) {
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
	user, err = model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "server error ..."
		}
	} else {
		loginResMes.Code = 200
		fmt.Println("user login success!")
	}

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
