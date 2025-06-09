package main

import (
	"fmt"
	"net"
)

// // 发送响应消息到客户端
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//依旧先发送长度
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buffer [4]byte
// 	binary.BigEndian.PutUint32(buffer[:], pkgLen)
// 	//发送
// 	_, err = conn.Write(buffer[:])
// 	if err != nil {
// 		fmt.Println("长度发送失败", err)
// 		return
// 	}

// 	//发送数据本身
// 	n, err := conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("响应消息发送失败", err)
// 		return
// 	}
// 	return
// }

// // 实现处理登录函数
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

// 	var loginmes message.LoginMes
// 	//获取mes的data并反序列化

// 	err = json.Unmarshal([]byte(mes.MetaData), &loginmes)
// 	if err != nil {
// 		fmt.Println("登录时反序列化失败", err)
// 		return
// 	}

// 	//接下来从数据库判断用户是否存在，并返消息
// 	var resmes message.Message
// 	resmes.Type = message.LoginRspType
// 	var loginrsp message.LoginRsp

// 	//判断账户
// 	if loginmes.UserId == 100 && loginmes.UserPwd == "123456" {
// 		loginrsp.Code = 200
// 	} else {
// 		loginrsp.Code = 500
// 		loginrsp.Error = "该账户未注册，请先注册！"
// 	}
// 	//将响应消息发送回客户端

// 	loginrspdata, err := json.Marshal(loginrsp)
// 	if err != nil {
// 		fmt.Println("登录时响应消息序列化失败", err)
// 		return
// 	}
// 	//先将loginrsp序列化装进resmes
// 	resmes.MetaData = string(loginrspdata)
// 	//再对resmes结构体进行序列化准备发送会客户端
// 	resmesdata, err := json.Marshal(resmes)
// 	if err != nil {
// 		fmt.Println("登录时响应消息序列化失败", err)
// 		return
// 	}
// 	//发送数据
// 	err = writePkg(conn, resmesdata)

// 	return
// }

// // 实现ServerProcessMes函数，根据消息种类不同，调用不同的函数
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 		//
// 	default:
// 		fmt.Println("消息类型不存在，无法处理！")

// 	}
// 	return
// }

// func readPkg(conn net.Conn) (mes message.Message, err error) {

// 	buffer := make([]byte, 8096)
// 	fmt.Println("Waiting for data from client")

// 	_, err = conn.Read(buffer[:4])
// 	if err != nil {
// 		// fmt.Println("Conn Read package length failed err is ", err)
// 		return
// 	}

// 	// fmt.Println("Read data is ", buffer[:4])
// 	//将读到的长度转化为uint32
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buffer[:4])
// 	//根据pkglen 从连接中读取数据
// 	n, err := conn.Read(buffer[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("Conn Read message failed err is ", err)
// 		return
// 	}
// 	//读取数据完成，将数据反序列化成message.Message类型
// 	err = json.Unmarshal(buffer[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("数据反序列化成message.Message类型失败，错误是", err)
// 		return
// 	}
// 	return
// }

// 处理和客户端的通信
func process(conn net.Conn) {
	//处理客户段发送的信息
	defer conn.Close()
	//调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程出现问题", err)
	}

}
func main() {
	//监听本地端口
	fmt.Println("新服务器结构在8080端口监听")
	listen, err := net.Listen("tcp", "10.10.4.137:8080")
	defer listen.Close()
	if err != nil {
		fmt.Println("Listen err=", err)
		return
	}
	defer listen.Close()

	//监听成功等待客户端连接服务器
	for {
		fmt.Println("Waiting for connection")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Failed to connection server err=", err)
		}
		//连接成功开始协程和客户端保持联系
		go process(conn)
	}
}
