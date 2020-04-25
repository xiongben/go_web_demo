package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_web_demo/chatdemo/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的信息")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 && err != nil {
		fmt.Println("conn write err = ", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) && err != nil {
		fmt.Println("conn write err = ", err)
		return
	}
	return
}
