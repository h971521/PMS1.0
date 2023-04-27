package operates

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"net"
)

func Register() (registerRes message.RegistrResMes, err error) {
	var userId int
	var userPwd string
	var roomNum int //房间数量
	fmt.Println("请输入用户名（手机号）")
	fmt.Scanln(&userId)
	fmt.Println("请输入密码")
	fmt.Scanln(&userPwd)
	fmt.Println("请输入房间数量")
	fmt.Scanln(&roomNum)

	//链接服务器
	conn, err := net.Dial("tcp", "192.168.56.1:8888")
	if err != nil {
		fmt.Println("register net.Dial fail err=", err)
		return
	}
	defer conn.Close()
	//向服务器发送注册信息
	var registerMes message.RegisterMes
	registerMes.UserId = userId
	registerMes.UserPwd = userPwd
	registerMes.RoomNum = roomNum
	data, _ := json.Marshal(registerMes)
	var mes message.Message
	mes.Type = message.RegisterMesType
	mes.Data = string(data)
	data, _ = json.Marshal(mes)
	var writerMes tools.Temp
	writerMes.Conn = conn
	writerMes.WritePkg(data)

	//等待服务器反馈..........
	var tf tools.Temp
	tf.Conn = conn
	mesres, err := tf.ReadPkg()
	//var registerRes message.RegistrResMes
	err = json.Unmarshal([]byte(mesres.Data), &registerRes)
	return
}
