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
	rows, err := this.pool.Query("SELECT * FROM user where id = ?", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var user = UserSql{}
		rows.Scan(&user.Id, &user.Username, &user.Birthday, &user.Sex, &user.Address, &user.Password)
		fmt.Println(user)
	}
	return
}
