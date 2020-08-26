package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"html/template"
	"net/http"
)

func main() {
	//http.Handle("/rest/user/",user2.MakeMuxer("/rest/user/"))

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("temperlate/static/"))))
	//http.HandleFunc("/http", handler)
	//http.HandleFunc("/testjson", testJsonRes)
	//http.HandleFunc("/testtemp", testTemp)
	//http.ListenAndServe(":8089", nil)

	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Logger.Fatal(e.Start(":1323"))

	app := iris.New()
	app.Logger().SetLevel("debug")
	//（可选）添加两个内置处理程序
	//可以从任何与http相关的panics中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())
	// 请求方法: GET
	// 资源标识: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome,xiongben!</h1>")
	})
	// 等同于 app.Handle("GET", "/ping", [...])
	// 请求方法: GET
	// 资源标识: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	// 请求方法: GET
	// 资源标识: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8086"), iris.WithoutServerError(iris.ErrServerClosed))
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
