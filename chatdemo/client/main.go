package main

import "fmt"

func main() {
	fmt.Println("this is client page")
	var myBuf [8096]byte
	fmt.Println(myBuf[:4])
}
