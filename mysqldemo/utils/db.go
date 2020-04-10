package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:xb2010550918@tcp(localhost:3306)/spring")
	if err != nil {
		panic(err.Error())
	}

}
