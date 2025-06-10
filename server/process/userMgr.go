package process2

import "fmt"

var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 初始化userMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成olineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 完成olineUsers删除
func (this *UserMgr) DeOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据Id返回对应的Userprocess
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d当前不在线", userId)
		return
	}
	return
}
