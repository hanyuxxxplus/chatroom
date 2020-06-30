package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8069]byte // 传输时用到的缓冲
}

func (this *Transfer)ReadPkg()(mes message.Message , err error){
	fmt.Println("读取客户端发送的数据")
	// conn.Read 在conn没有被关闭的情况才会被阻塞
	// 如果客户关闭了连接 就不会阻塞
	_ , err = this.Conn.Read(this.Buf[:4])

	if err != nil {
		fmt.Println("read err = ",err)
		return mes,err
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n , err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.read err = ",err)
		return mes,err
	}
	// 将buf反序列化成message.Message
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ",err)
		return mes,err
	}
	return mes,err
}

// 封装 writePkg
func (this *Transfer)WritePkg(data []byte)(err error){
	// 发送一个长度给对方，因为是在一次连接中，所以先发了一次长度之后 连接并没有断开，发送玩之后 再发送数据
	pkgLen := uint32(len(data))
	dataBytes := make([]byte,4)
	binary.BigEndian.PutUint32(dataBytes,pkgLen)
	n , err := this.Conn.Write(dataBytes)
	if n != 4 || err != nil {
		fmt.Println("conn.Write() err = ",err)
		return
	}
	n , err = this.Conn.Write(data)
	if  uint32(n) !=  pkgLen || err != nil {
		fmt.Println("conn.Write() err = ",err)
		return
	}
	return
}