package model

import (
	"database/sql"
	"fmt"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *sql.DB
}

func NewUserDao(pool *sql.DB) (userDao *UserDao) {
	userDao = &UserDao{pool: pool}
	return
}

func (this *UserDao) GetUserById(id int) (user UserSql, err error) {
	var name string
	var password string
	err = this.pool.QueryRow("SELECT username,password FROM user where id = ?", id).Scan(&name, &password)
	if err == sql.ErrNoRows {
		fmt.Println("there is no user's id is XXXX")
		err = ERROR_USER_NOTEXISTS
		return
	} else {
		user.Username = name
		user.Password = password
		fmt.Println(user)
	}
	//for rows.Next() {
	//	var user = UserSql{}
	//	rows.Scan(&user.Id, &user.Username, &user.Birthday, &user.Sex, &user.Address, &user.Password)
	//	fmt.Println(user)
	//}
	return
}

func (this *UserDao) Login(userid int, userPwd string) (user UserSql, err error) {
	user, err = this.GetUserById(userid)
	if err != nil {
		return
	}
	if user.Password != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
