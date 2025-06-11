package model

import (
	"MessageSystem/common/message"
	"net"
)

// 客户端会使用很多
type CurUser struct {
	Conn net.Conn
	message.User
}
