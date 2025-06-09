package process2

import (
	"MessageSystem/common/message"
	"MessageSystem/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

// 实现处理登录函数
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	var loginmes message.LoginMes
	//获取mes的data并反序列化

	err = json.Unmarshal([]byte(mes.MetaData), &loginmes)
	if err != nil {
		fmt.Println("登录时反序列化失败", err)
		return
	}

	//接下来从数据库判断用户是否存在，并返消息
	var resmes message.Message
	resmes.Type = message.LoginRspType
	var loginrsp message.LoginRsp

	//判断账户
	if loginmes.UserId == 100 && loginmes.UserPwd == "123456" {
		loginrsp.Code = 200
	} else {
		loginrsp.Code = 500
		loginrsp.Error = "该账户未注册，请先注册！"
	}
	//将响应消息发送回客户端

	loginrspdata, err := json.Marshal(loginrsp)
	if err != nil {
		fmt.Println("登录时响应消息序列化失败", err)
		return
	}
	//先将loginrsp序列化装进resmes
	resmes.MetaData = string(loginrspdata)
	//再对resmes结构体进行序列化准备发送会客户端
	resmesdata, err := json.Marshal(resmes)
	if err != nil {
		fmt.Println("登录时响应消息序列化失败", err)
		return
	}
	//发送数据
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(resmesdata)

	return
}
