package main

import (
	"PMS/message"
	"PMS/sever/process"
	"PMS/tools"
	"fmt"
	"net"
)

func main() {
	//0.0.0.0：表示在本地监听 8888 端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	fmt.Println("----------正在监听网络端口8888----------")

	for {
		conn, err := listen.Accept() //获取端口conn
		if err != nil {
			fmt.Println("端口获取失败！")
			return
		}
		defer conn.Close()
		var getMes tools.Temp
		getMes.Conn = conn
		var mes message.Message
		mes, err = getMes.ReadPkg()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(mes.Type)
		switch mes.Type {
		case message.LoginMesType:
			process.LoginRes(mes.Data, conn)
		case message.RegisterMesType:
			process.RegisterRes(mes.Data, conn)
		case message.HomePageType:
			process.HomePageRes(mes.Data, conn)
		case message.DayRoomMes:
			process.RoomStartState(mes.Data, conn)
		case message.AddOrdersType:
			process.AddOrderRes(mes.Data, conn)
		case message.DelOrdersType:
			process.DelOrderRes(mes.Data, conn)
		case message.ReadOrderType:
			process.ReadOrderRes(mes.Data, conn)
		case message.TurnOrderType:
			process.TurnOrderRes(mes.Data, conn)
		}
	}
}
