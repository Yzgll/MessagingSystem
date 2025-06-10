package message

const (
	LoginMesType            = "LoginMes"
	LoginRspType            = "LoginRsp"
	RegisterMesType         = "RegisterMes"
	RegisterRspType         = "RegisterRsp"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

// 用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type     string `json:"type"` //消息类型
	MetaData string `json:"metaData"`
}

// 登录发送消息
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户名
	UserPwd  string `json:"userPwd"`  //密码
	UserName string `json:"userName"` //用户名
}
type LoginRsp struct {
	Code    int    `json:"code"` //状态码500表示未注册 200表示登录成功
	UsersId []int  `json:"usersId"`
	Error   string `json:"error"` //错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}
type RegisterRsp struct {
	Code  int    `json:"code"`  //状态码505表示已经占用 200表示登录成功
	Error string `json:"error"` //错误信息
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
