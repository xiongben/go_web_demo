package clientprocess

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_web_demo/chatdemo/common/message"
	"go_web_demo/chatdemo/server/utils"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial err =", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.marshal err = ", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err = ", err)
		return
	}
	//先发送数据长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write byte err =", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write byte err =", err)
		return
	}

	//接收服务器返回消息部分逻辑
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("conn.read byte err =", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	fmt.Println(loginResMes)
	return
}
