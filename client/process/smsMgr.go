package process

import (
	"MessageSystem/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var smsmes message.SmsMes
	err := json.Unmarshal([]byte(mes.MetaData), &smsmes)
	if err != nil {
		fmt.Println("outputGroupMes反序列化失败 err is ", err)
		return
	}
	info := fmt.Sprintf("用户Id:\t%d 发送消息:\t%s", smsmes.UserId, smsmes.Content)
	fmt.Println(info)
	fmt.Println()
}
