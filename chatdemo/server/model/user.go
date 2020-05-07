package model

type User struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	UserPwd  string `json:"user_pwd"`
}

type UserSql struct {
	Id       int
	Username string
	Birthday string
	Sex      string
	Address  string
	Password string
}

type Testss struct {
	Name string
}
