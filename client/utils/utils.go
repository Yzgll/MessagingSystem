package utils

import (
	"MessageSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 分层加面向对象，将方法封装到结构体
type Transfer struct {
	Conn   net.Conn
	Buffer [8096]byte //缓冲区
}

// 发送响应消息到客户端
func (this *Transfer) WritePkg(data []byte) (err error) {
	//依旧先发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buffer [4]byte
	binary.BigEndian.PutUint32(this.Buffer[0:4], pkgLen)
	//发送
	n, err := this.Conn.Write(this.Buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("长度发送失败", err)
		return
	}

	//发送数据本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("消息发送失败", err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buffer := make([]byte, 8096)
	fmt.Println("Waiting for data from client")

	_, err = this.Conn.Read(this.Buffer[:4])
	if err != nil {
		// fmt.Println("Conn Read package length failed err is ", err)
		return
	}

	// fmt.Println("Read data is ", buffer[:4])
	//将读到的长度转化为uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buffer[0:4])
	//根据pkglen 从连接中读取数据
	n, err := this.Conn.Read(this.Buffer[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("Conn Read message failed err is ", err)
		return
	}
	//读取数据完成，将数据反序列化成message.Message类型
	err = json.Unmarshal(this.Buffer[:pkgLen], &mes)
	if err != nil {
		fmt.Println("数据反序列化成message.Message类型失败，错误是", err)
		return
	}
	return
}
