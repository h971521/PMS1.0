package operates

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"net"
)

// 登录操作 --> 向服务器发送用户 ID 和 密码 确认该用户是否存在
func Login(userId int, userPwd string) (loginmes message.LoginResMes, err error) {
	//链接服务器
	conn, err := net.Dial("tcp", "192.168.56.1:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//向服务器发送登录信息
	var userMes message.UserMes
	userMes.UserId = userId
	userMes.UserPwd = userPwd
	//将用户信息序列化
	data, err := json.Marshal(userMes)
	if err != nil {
		fmt.Println("login json.Marshal(userMes) fail err=", err)
		return
	}
	var loginMes message.Message         //准备发送给服务器的数据
	loginMes.Type = message.LoginMesType //消息类型
	loginMes.Data = string(data)
	//将其序列化后发送给服务器
	data, err = json.Marshal(loginMes)
	if err != nil {
		fmt.Println("login json.Marshal(loginMes) fail err=", err)
		return
	}

	//此时data就是要发送的消息
	var loginMesWrite tools.Temp
	loginMesWrite.Conn = conn
	err = loginMesWrite.WritePkg(data) //向服务器发送信息
	if err != nil {
		fmt.Println("loginMesWrite fail err=", err)
		return
	}
	fmt.Println("登录信息发送成功")
	//等待服务器的反馈
	var loginMesRead tools.Temp
	loginMesRead.Conn = conn
	var loginResMes message.Message
	loginResMes, err = loginMesRead.ReadPkg()
	if err != nil {
		fmt.Println("loginMesRead.ReadPkg() fail err=", err)
		return
	}
	err = json.Unmarshal([]byte(loginResMes.Data), &loginmes)
	return
}
