package process

import (
	"MessageSystem/client/utils"
	"MessageSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

// 登录函数
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//1.连接到服务器
	conn, errDial := net.Dial("tcp", "10.10.4.137:8080")
	if errDial != nil {
		fmt.Println("Failed to Dial err is ", errDial)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器,创建消息结构体
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建LoginMes结构体
	var loginmes message.LoginMes
	loginmes.UserId = userId
	loginmes.UserPwd = userPwd
	//4.把消息序列化装到mes结构体的MeteData中
	data, errMarshal := json.Marshal(loginmes)
	if errMarshal != nil {
		fmt.Println("Json Marshal failed err is ", errMarshal)
	}
	//5.装到mes结构体的MeteData中
	mes.MetaData = string(data)
	//6.再把mes结构体序列化，用于发送到服务器
	mesdata, errmesMarshal := json.Marshal(mes)
	if errmesMarshal != nil {
		fmt.Println("Mes Json Marshal failed err is ", errmesMarshal)
	}

	//7.发送到服务器
	//7.1先发送数据的长度
	//将数据长度转化为切片
	var dataLen uint32
	dataLen = uint32(len(mesdata))
	var buffer [4]byte
	binary.BigEndian.PutUint32(buffer[:], dataLen)
	fmt.Println(buffer)
	n, err := conn.Write(buffer[:])
	if n != 4 || err != nil {
		fmt.Println("Conn Write err is ", err)
	}
	fmt.Printf("The length of mesdata is %d and contex is%s \n", len(mesdata), string(mesdata))

	//发送消息本身
	_, err = conn.Write(mesdata)
	if err != nil {
		fmt.Println("Conn Write mesdata err is ", err)
		return
	}
	//处理服务器返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("读取服务器返回消息失败", err)
		return
	}

	var loginrsp message.LoginRsp
	err = json.Unmarshal([]byte(mes.MetaData), &loginrsp)
	if loginrsp.Code == 200 {
		go serverProcessMes(conn)
		//登录成功，显示登录成功的菜单
		for {
			ShowMenu()
		}
		//启动协程保持和服务器的通讯，如果有服务器推送的消息则显示在客户端的终端

	} else {
		fmt.Println(loginrsp.Error)
	}
	return
}
