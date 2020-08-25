package main

import (
	"fmt"
	"go_web_demo/chatdemo/server/model"
	"go_web_demo/chatdemo/server/process"
	"net"
)

func init() {
	process.InitPool()
	initUserDao()
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(process.Pool)
}

func main() {
	//model.MyUserDao.GetUserById(55)
	fmt.Println("server is listen in 8889,,,,,")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net listen error = ", err)
		return
	}
	for {
		fmt.Println("waiting client to connct , , ,")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connect accept error = ", err)
			return
		}
		go processfn(conn)

	}
}

func processfn(conn net.Conn) {
	defer conn.Close()
	processor1 := &process.ProcessorObj{Conn: conn}
	err := processor1.SeverProcess()
	if err != nil {
		fmt.Println("the connect between client and server have some err = ", err)
		return
	}
}
