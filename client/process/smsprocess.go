package process

import (
	"MessageSystem/client/utils"
	"MessageSystem/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// 发送群聊消息函数
func (this *SmsProcess) SendGroupMes(connet string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsmes message.SmsMes
	smsmes.Content = connet
	smsmes.User = CurUser.User
	data, err := json.Marshal(smsmes)
	if err != nil {
		fmt.Println("SendGroupMes序列化失败", err)
	}
	mes.MetaData = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes序列化失败", err)
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes", err)
	}
	return
}
