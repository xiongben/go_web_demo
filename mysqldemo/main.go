package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

type User struct {
	id       int
	username string
	birthday string
	sex      string
	address  string
}

func main() {
	db, err := sql.Open("mysql", "root:xb2010550918@tcp(localhost:3306)/spring")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}
	//columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for rows.Next() {
		var student User
		err = rows.Scan(&student.id, &student.username, &student.birthday, &student.sex, &student.address)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println(student)
	}
}
