package operates

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// 添加订单
func AddOrder(mes message.RoomUpdate) (addmes message.TotalMes) {
	conn, _ := net.Dial("tcp", "192.168.56.1:8888")
	defer conn.Close()
	var tf tools.Temp
	tf.Conn = conn //链接服务器
	for {
		fmt.Println("请选择您要进行的操作")

		fmt.Println("1.客人姓名：")

		fmt.Println("2.客人电话：")

		fmt.Println("3.房间价格：")

		fmt.Println("0.返回上一级")
		var key int //从命令行接收要进行的操作
		fmt.Scanln(&key)
		switch key {
		case 0:
			//计算当日总房价
			if mes.CustomersName == "" && mes.CustomersTel == 0 { //当对该房间没有任何操作时，直接输出
				data, err := json.Marshal(mes)
				if err != nil {
					fmt.Println("roomupdate json fail err=", err)
					return
				}
				var mesage message.Message
				mesage.Type = message.AddOrdersType
				mesage.Data = string(data)
				data, err = json.Marshal(mesage)
				if err != nil {
					fmt.Println("json roomupdatemes fail err=", err)
					return
				}
				//向服务器发送添加的订单
				err = tf.WritePkg(data)
				if err != nil {
					fmt.Println("send to serve roomupdatemes fail err=", err)
					return
				}
				//等待服务器反馈房态信息

				addres, _ := tf.ReadPkg()
				addresdata := addres.Data
				json.Unmarshal([]byte(addresdata), &addmes)
				return
				return
			}
			//获取订单录入时间
			t := time.Now()
			hour := t.Hour()
			minute := t.Minute()
			tim := fmt.Sprintf("%v:%v", hour, minute) //订单办理时间
			//time为该订单的录如时间，将其赋给房间信息
			mes.CheckInTime = tim
			data, err := json.Marshal(mes)
			if err != nil {
				fmt.Println("roomupdate json fail err=", err)
				return
			}

			var mesage message.Message
			mesage.Type = message.AddOrdersType
			mesage.Data = string(data)
			data, err = json.Marshal(mesage)
			if err != nil {
				fmt.Println("json roomupdatemes fail err=", err)
				return
			}
			//向服务器发送添加的订单
			err = tf.WritePkg(data)
			if err != nil {
				fmt.Println("send to serve roomupdatemes fail err=", err)
				return
			}
			//等待服务器反馈房态信息

			addres, err := tf.ReadPkg()
			addresdata := addres.Data
			//addmes为服务器返回的消息
			json.Unmarshal([]byte(addresdata), &addmes)
			return
		case 1:
			var custName string //客人姓名
			fmt.Scanln(&custName)
			mes.CustomersName = custName
			continue
		case 2:
			var custTel int //客人电话
			fmt.Scanln(&custTel)
			mes.CustomersTel = custTel
			continue
		case 3:
			var price float32 //房间价格
			fmt.Scanln(&price)
			mes.Price = price
		default:
			fmt.Println("输入有误请重新输入")
			continue
		}
	}
}
func DelOrder(mes message.RoomUpdate) (res string) {
	//Roomnum int //房间号
	//RoomMes
	//Update{Day、UserId}

	//链接服务器
	conn, _ := net.Dial("tcp", "192.168.56.1:8888")
	defer conn.Close()
	var tf tools.Temp
	tf.Conn = conn
	data, _ := json.Marshal(mes)
	var delMes message.Message
	delMes.Type = message.DelOrdersType
	delMes.Data = string(data)
	data, _ = json.Marshal(delMes)
	//向服务器发送要删除的信息
	tf.WritePkg(data)
	//等待服务器反馈
	delRes, _ := tf.ReadPkg()
	res = delRes.Data
	return
}

// 查看指定日期指定房间的房间信息
func ReadOrder(mes message.RoomUpdate) (readres message.RoomMes) {
	//链接服务器
	conn, _ := net.Dial("tcp", "192.168.56.1:8888")
	defer conn.Close()
	var tf tools.Temp
	tf.Conn = conn
	data, _ := json.Marshal(mes)
	var delMes message.Message
	delMes.Type = message.ReadOrderType
	delMes.Data = string(data)
	data, _ = json.Marshal(delMes)
	//向服务器发送要查看的信息
	tf.WritePkg(data)
	//等待服务器反馈
	readRes, _ := tf.ReadPkg()
	json.Unmarshal([]byte(readRes.Data), &readres)
	return
}

// 修改指定订单
func TurnOrder(mes message.RoomUpdate) (turnmes message.TotalMes) {
	//链接服务器
	conn, _ := net.Dial("tcp", "192.168.56.1:8888")
	defer conn.Close()
	var tf tools.Temp
	tf.Conn = conn
	for {
		fmt.Println("请选择要修改的内容")
		fmt.Println("1.修改客人姓名")
		fmt.Println("2.修改客人电话")
		fmt.Println("3.修改房间价格")
		fmt.Println("0.确认修改并提出")
		var key int
		fmt.Scanln(&key)
		switch key {
		case 1:
			var turnName string
			fmt.Scanln(&turnName)
			mes.CustomersName = turnName
			continue
		case 2:
			var turnTel int
			fmt.Scanln(&turnTel)
			mes.CustomersTel = turnTel
			continue
		case 3:
			var turnPrice float32
			fmt.Scanln(&turnPrice)
			mes.Price = turnPrice
			continue
		case 0:
			if mes.CustomersTel == 0 && mes.CustomersName == "" && mes.Price == 0 {
				return
			} else {
				data, _ := json.Marshal(mes)
				var turnOrderMes message.Message
				turnOrderMes.Type = message.TurnOrderType
				turnOrderMes.Data = string(data)
				data, _ = json.Marshal(turnOrderMes)
				//向服务器发送要修改的信息
				tf.WritePkg(data)
				fmt.Println("修改信息发送成功")
				//等待服务器反馈
				turnres, _ := tf.ReadPkg()
				turnresdata := turnres.Data
				//addmes为服务器返回的消息
				json.Unmarshal([]byte(turnresdata), &turnmes)
				fmt.Println("修改信息读取成功")
				return
			}
		default:
			fmt.Println("输入有误请重新输入")
			continue
		}
	}
}
