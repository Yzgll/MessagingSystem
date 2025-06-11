package process2

import (
	"MessageSystem/client/utils"
	"MessageSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//群发遍历onlineUsers把消息群发出去
	var smsmes message.SmsMes
	err := json.Unmarshal([]byte(mes.MetaData), &smsmes)
	if err != nil {
		fmt.Println("SendGroupMes Unmarshal Err is", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes Marshal Err is", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//不发给自己
		if id == smsmes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}
