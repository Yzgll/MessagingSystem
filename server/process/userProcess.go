package process2

import (
	"MessageSystem/common/message"
	"MessageSystem/server/model"
	"MessageSystem/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int //表示是哪个用户的连接
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

	//判断账户现在需要到数据库去判断
	user, err := model.MyUserDao.Login(loginmes.UserId, loginmes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXITS {
			loginrsp.Code = 500
			loginrsp.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginrsp.Code = 403
			loginrsp.Error = err.Error()
		} else {
			loginrsp.Code = 505
			loginrsp.Error = "服务器内部错误"
		}

	} else {
		loginrsp.Code = 200
		//登录成功，将该用户放入在线列表中
		this.UserId = loginmes.UserId
		userMgr.AddOnlineUser(this)
		//登录成功后通知其他用户我已经上线
		this.NotifyOthersOnlineUser(loginmes.UserId)
		//将在在线用户列表的所有用户放入新增的切片返回
		for id, _ := range userMgr.onlineUsers {
			loginrsp.UsersId = append(loginrsp.UsersId, id)
		}
		fmt.Println(user.UserId, "登录成功")
	}

	// if loginmes.UserId == 100 && loginmes.UserPwd == "123456" {
	// 	loginrsp.Code = 200
	// } else {
	// 	loginrsp.Code = 500
	// 	loginrsp.Error = "该账户未注册，请先注册！"
	// }
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

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//反序列化获取user结构体
	var registermes message.RegisterMes
	err = json.Unmarshal([]byte(mes.MetaData), &registermes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	//注册成功返回响应消息
	var resmes message.Message
	resmes.Type = message.RegisterRspType

	var registerrsp message.RegisterRsp
	//调用userdao的注册方法
	err = model.MyUserDao.Register(&registermes.User)

	if err != nil {
		if err == model.ERROR_USER_EXITS {
			registerrsp.Code = 505
			registerrsp.Error = model.ERROR_USER_EXITS.Error()
		} else {
			registerrsp.Code = 506
			registerrsp.Error = "未知错误"
		}
	} else {
		registerrsp.Code = 200
	}

	//将响应消息返回
	data, err := json.Marshal(registerrsp)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	resmes.MetaData = string(data)

	data, err = json.Marshal(resmes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	//发送数据
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}

// userId通知其他在线用户，自己上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历维护的onlineusersmap
	//userId这是自己的是当前刚刚上线的
	for id, up := range userMgr.onlineUsers {
		//不通知自己
		if id == userId {
			continue
		}
		up.NotifyMeToOthersOnline(userId)
	}
}

func (this *UserProcess) NotifyMeToOthersOnline(userId int) {
	//发送NotifyUserStatusMes类型的消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes.MetaData = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	//发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeToOthersOnline", err)
		return
	}
}
