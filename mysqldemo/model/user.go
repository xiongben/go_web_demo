package model

import (
	"fmt"
	"go_web_demo/mysqldemo/utils"
)

type User struct {
	id       int
	username string
	birthday string
	sex      string
	address  string
}

//预编译
func (user *User) AddUser() error {
	sqlStr := "insert into user(username,birthday,sex,address) values(?,?,?,?)"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
	}
	_, err2 := inStmt.Exec("日向镜", "1993-03-02", "男", "木叶村")
	if err2 != nil {
		fmt.Println("执行sql语句出现异常：", err2)
		return err2
	}
	return nil
}

func (user *User) AddUser2() error {
	sqlStr := "insert into user(username,birthday,sex,address) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, "日向白", "1993-03-02", "男", "木叶村")
	defer utils.Db.Close()
	if err != nil {
		fmt.Println("执行sql出现异常", err)
		return err
	}
	return nil
}

func (user *User) FindUserInfo() ([]User, error) {
	sqlStr := "SELECT * FROM user"
	rows, err := utils.Db.Query(sqlStr)
	defer utils.Db.Close()
	if err != nil {
		fmt.Println("执行sql出现异常", err)
		return nil, err
	}
	var users []User
	for rows.Next() {
		var student User
		err = rows.Scan(&student.id, &student.username, &student.birthday, &student.sex, &student.address)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		users = append(users, student)
	}
	return users, nil
}

func (user *User) GetUserById() (*User, error) {
	sqlStr := "SELECT * FROM user where id = ?"
	row := utils.Db.QueryRow(sqlStr, user.id)
	defer utils.Db.Close()

	var student User
	err := row.Scan(&student.id, &student.username, &student.birthday, &student.sex, &student.address)
	if err != nil {
		fmt.Println("执行sql出现异常", err)
		return nil, err
	}
	return &student, nil
}
