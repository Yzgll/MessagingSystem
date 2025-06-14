package message

const (
	LoginMesType    = "LoginMes"
	LoginRspType    = "LoginRsp"
	RegisterMesType = "RegisterMes"
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
	Code  int    `json:"code"`  //状态码500表示未注册 200表示登录成功
	Error string `json:"error"` //错误信息
}

type RegisterMes struct {
}
