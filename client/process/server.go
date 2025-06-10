package process

import (
	"MessageSystem/client/utils"
	"MessageSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("**************************************")
	fmt.Println("*          恭喜XXX登录成功           *")
	fmt.Println("*                                    *")
	fmt.Println("* 1. 显示用户在线列表                *")
	fmt.Println("* 2. 发送消息                        *")
	fmt.Println("* 3. 信息列表                        *")
	fmt.Println("* 4. 退出系统                        *")
	fmt.Println("*                                    *")
	fmt.Println("* 请选择(1-4):                       *")
	fmt.Println("**************************************")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUsers()
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("操作有误...")
	}
}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//读取服务器的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器推送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("读取服务器推送消息失败错误是", err)
			return
		}
		//正确读取到消息
		//fmt.Printf("正确读取到消息%v\n", mes)
		//处理服务器发来的消息
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//取出消息
			//加入到客户端维护的map
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.MetaData), &notifyUserStatusMes)
			updataUserStatus(&notifyUserStatusMes)
		default:
			fmt.Println("服务器返回了未知消息类型")
		}
	}
}
