package process

import (
	"database/sql"
	"fmt"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

type user struct {
	id       int
	username string
	birthday string
	sex      string
	address  string
	password string
}

var Pool *sql.DB

func InitPool() (err error) {
	Pool, err = sql.Open("mysql", "root:xb2010550918@/spring")
	//defer Pool.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Testmysql() {
	pool, err := sql.Open("mysql", "root:xb2010550918@/spring")
	defer pool.Close()

	rows, err := pool.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user = user{}
		rows.Scan(&user.id, &user.username, &user.birthday, &user.sex, &user.address, &user.password)
		fmt.Println(user)
	}

}
