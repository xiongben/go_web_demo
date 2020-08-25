package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//http.Handle("/rest/user/",user2.MakeMuxer("/rest/user/"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("temperlate/static/"))))
	http.HandleFunc("/http", handler)
	http.HandleFunc("/testjson", testJsonRes)
	http.HandleFunc("/testtemp", testTemp)
	http.ListenAndServe(":8089", nil)

	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Logger.Fatal(e.Start(":1323"))
}

func testTemp(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("temperlate/a.html"))
	t.Execute(w, "")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "测试http协议", r.URL.Path)
	fmt.Fprintln(w, "测试http协议", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的信息", r.Header.Get("Accept-Encoding"))
	fmt.Fprintln(w, "请求头中的信息", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中的信息", r.Header)

	//len := r.ContentLength
	//body := make([]byte, len)
	//r.Body.Read(body)
	//fmt.Fprintln(w,"请求体中的内容有:", string(body))

	pass := r.PostFormValue("pass")
	fmt.Fprintln(w, "pass: ", pass)
	//username := r.PostFormValue("username")
	//fmt.Fprintln(w,"username: ", username)
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	cat := Animal{
		Name: "cat",
		Age:  2,
	}
	json, _ := json.Marshal(cat)
	w.Write(json)
}

type Animal struct {
	Name string
	Age  int
}
