package process

import (
	"encoding/json"
	"fmt"
	"go_web_demo/chatdemo/common/message"
	"go_web_demo/chatdemo/server/utils"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err = ", err)
		return
	}
	for id, up := range userMgr.onlineusers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("send message to each online user error = ", err)
	}

}
