package process

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineusers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{onlineusers: make(map[int]*UserProcess, 1024)}
}

func (this *UserMgr) AddONlineUser(up *UserProcess) {
	this.onlineusers[up.UserId] = up
}

func (this *UserMgr) DelONlineUser(userId int) {
	delete(this.onlineusers, userId)
}

func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineusers
}

func (this *UserMgr) GetAllOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineusers[userId]
	if !ok {
		err = fmt.Errorf("user#{userId} not exist!")
		return
	}
	return
}
