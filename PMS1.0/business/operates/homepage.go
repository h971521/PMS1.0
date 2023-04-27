package operates

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func HomePage(userId int) (totalPrice float32, roomBalance int) {
	//var totalPrice float32 //今日总价
	//var roomBalance int    //剩余放量

	//链接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	datenow := fmt.Sprintf("%02d-%02d", month, day)
	var HPMes message.HomePage
	HPMes.UserId = userId
	HPMes.Day = datenow
	data, err := json.Marshal(HPMes)
	var mes message.Message
	mes.Type = message.HomePageType
	mes.Data = string(data)
	data, _ = json.Marshal(mes)
	var tf tools.Temp
	tf.Conn = conn
	tf.WritePkg(data)
	//等待服务器响应
	var hpResMes message.HomePageRes
	datares, _ := tf.ReadPkg()
	json.Unmarshal([]byte(datares.Data), &hpResMes)
	roomBalance = hpResMes.RoomBalance
	totalPrice = hpResMes.TotalPrice
	fmt.Printf("\t\t\t\t%v-%02d-%02d\n", year, month, day)
	return
}
