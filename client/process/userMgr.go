package process

import (
	"MessageSystem/common/message"
	"fmt"
)

// 客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUsers()
}

// 在客户端显示当前在线的yoghurt
func outputOnlineUsers() {
	fmt.Println("当前用户在线列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户Id:\t", id)
	}
}
