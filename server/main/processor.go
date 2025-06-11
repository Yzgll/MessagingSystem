package main

import (
	"MessageSystem/common/message"
	process2 "MessageSystem/server/process"
	"MessageSystem/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 实现ServerProcessMes函数，根据消息种类不同，调用不同的函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//测试是否能收到群发的消息
	fmt.Printf("mes=\n", mes)
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册消息
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsprocess := &process2.SmsProcess{}
		smsprocess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理！")

	}
	return
}

func (this *Processor) process2() (err error) {
	for {

		//将读取数据包封装成函数，返回message
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return err
			} else {
				fmt.Println("读取包信息失败，错误是", err)
				return err
			}

		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}

	}
}
