package tools

import (
	"PMS/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Temp struct {
	Conn net.Conn
	Buf  [6048]byte
}

func (t *Temp) WritePkg(data []byte) (err error) {

	var pkglen uint32 //准备发送数据的长度
	pkglen = uint32(len(data))
	//将长度以转换为 []byte 类型
	binary.BigEndian.PutUint32(t.Buf[:4], pkglen)
	//向服务器\客户端发送数据的长度
	_, err = t.Conn.Write(t.Buf[:4])
	if err != nil {
		fmt.Println("data length conn.Write fail err=", err)
		return
	}
	n, err := t.Conn.Write(data)
	if n != len(data) || err != nil {
		fmt.Println("data conn.Write fail err=", err)
		return
	}
	return
}
func (t *Temp) ReadPkg() (mes message.Message, err error) {

	//读取数据长度
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		fmt.Println("Conn.Read data length fail err=", err)
		return
	}
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(t.Buf[:4])
	//读取信息
	n, err := t.Conn.Read(t.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("Conn.Read data fail err=", err)
		return
	}
	err = json.Unmarshal(t.Buf[:pkglen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal Message fail err=", err)
		return
	}
	return
}
