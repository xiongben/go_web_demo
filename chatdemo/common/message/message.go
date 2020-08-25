package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMesType"
	RegisterResMesType      = "RegisterResMesType"
	NotifyUserStatusMesType = "NotifyUserStatusMesType"
	SmsMesType              = "SmsMesType"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginResMes struct {
	Code    int    `json:"code"`
	UserIds []int  `json:"user_ids"`
	Error   string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`
	User
}
