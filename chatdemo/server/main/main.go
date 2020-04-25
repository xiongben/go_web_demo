package main

import (
	"fmt"
	"net"
)

func main() {
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

	}
}
