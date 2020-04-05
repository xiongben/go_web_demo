package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/http", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "测试http协议")
}
