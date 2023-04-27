package operates

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"net"
)

func RoomState(userId int) {
	var roomupdate message.RoomUpdate //要操作的房间信息（包括日期、房间号）
	//选择要操作的日期
	fmt.Println("请输入要操作的日期（ex.四月一日：04-01）")
	var day string
	fmt.Scanln(&day)
	//向服务器发送要操作的日期
	var mes message.Update
	mes.Day = day
	mes.UserId = userId

	sendmes, _ := json.Marshal(mes)
	var dayMes message.Message
	dayMes.Type = message.DayRoomMes
	dayMes.Data = string(sendmes)
	data, _ := json.Marshal(dayMes)

	//链接服务器
	conn, err := net.Dial("tcp", "192.168.56.1:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	var temp tools.Temp
	temp.Conn = conn
	temp.WritePkg(data) //向服务器发送要操作的日期

	//等待服务器反馈

	rdata, _ := temp.ReadPkg() //获取房间状态
	var dayroommes []message.RoomMes
	err = json.Unmarshal([]byte(rdata.Data), &dayroommes)
	if err != nil {
		fmt.Println(err)
	}
	roomnums := len(dayroommes)
	for {
		//输出当日所有房间信息
		for i := 0; i < roomnums; i++ {
			fmt.Printf("%+v\n", dayroommes[i])
		}
		fmt.Println()

		var roomnum int //要操作的房间号
		for {
			fmt.Println("(0.返回上一级)")
			fmt.Printf("您有%d间房间，请输入您要操作的房间号（1-%d）\n", roomnums, roomnums)
			fmt.Scanln(&roomnum)
			if roomnum > roomnums || roomnum < 0 {
				fmt.Println("房间号输入错误请重新输入！")
			} else {
				break
			}
		}
		//返回上一级（首页）
		if roomnum == 0 {
			break
		}
		//房间订单的增删查改操作
		roomupdate.UserId = userId
		roomupdate.Day = day
		roomupdate.Roomnum = roomnum - 1
		for {
			fmt.Println("请选择您要进行的操作")
			fmt.Println("1.添加订单")
			fmt.Println("2.修改订单")
			fmt.Println("3.删除订单")
			fmt.Println("4.查看订单")
			fmt.Println("5.返回上一级")
			var choice int
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				addorderres := AddOrder(roomupdate) //添加订单
				fmt.Println()
				fmt.Printf("今日房间余量:%v\n", addorderres.RoomBalance)
				//各房间状态
				for i := 0; i < len(addorderres.Rooms); i++ {
					fmt.Printf("%+v\n", addorderres.Rooms[i])
				}
				continue
			case 2:
				turnorderres := TurnOrder(roomupdate)
				fmt.Println()
				fmt.Printf("今日房间余量:%v\n", turnorderres.RoomBalance)
				//各房间状态
				for i := 0; i < len(turnorderres.Rooms); i++ {
					fmt.Printf("%+v\n", turnorderres.Rooms[i])
				}
				continue
			case 3:
				res := DelOrder(roomupdate)
				if res == "Success" {
					fmt.Println("删除订单成功")
				} else {
					fmt.Println("删除订单失败")
				}
				for {
					fmt.Println("0.返回上一级")
					var oper int
					fmt.Scanln(&oper)
					if oper == 0 {
						break
					} else {
						fmt.Println("输入错误请重新输入")
					}
				}
				continue
			case 4:
				readres := ReadOrder(roomupdate)
				fmt.Printf("房间%v的信息为：%+v\n", roomnum, readres)
				for {
					fmt.Println("0.返回上一级")
					var oper int
					fmt.Scanln(&oper)
					if oper == 0 {
						break
					} else {
						fmt.Println("输入错误请重新输入")
					}
				}
				continue
			case 5:
				break
			default:
				fmt.Println("输入有误请重新输入")
				continue
			}
			return
		}
	}
}
