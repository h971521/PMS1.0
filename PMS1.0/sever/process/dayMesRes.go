package process

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
)

// 读取某日房间状态
func RoomStartState(mesint string, c net.Conn) {
	var mes message.Update
	json.Unmarshal([]byte(mesint), &mes)
	//链接 redis 数据库
	conn := tools.GetRedisPool() //链接池获取数据库链接
	defer conn.Close()
	//获取房间数量
	roomnum, err := redis.Int(conn.Do("hget", mes.UserId, "roomnum"))
	if err != nil {
		fmt.Println("房量信息读取失败,err=", err)
		return
	}
	//初始化当日房间状态，如果不存在则向数据库写入当日房间状态
	var totleMes message.TotalMes //包括每个房间状态、今日总价、房间余量
	judge, err := redis.Bool(conn.Do("hexists", mes.UserId, mes.Day))
	//fmt.Println(judge)
	if judge == false {
		//向数据库发送容量为房间数量的切片、房间余量、今日总价
		var rooms = make([]message.RoomMes, roomnum)
		for i := 0; i < roomnum; i++ {
			var roommes message.RoomMes
			//rooms = append(rooms, roommes)
			rooms[i] = roommes
		}
		totleMes.Rooms = rooms
		totleMes.RoomBalance = roomnum
		totleMes.TotalPrice = 0
		//fmt.Printf("%T,%v", rooms, rooms)
		datatemp, _ := json.Marshal(totleMes)
		//fmt.Println(datatemp)
		_, err = conn.Do("hset", mes.UserId, mes.Day, datatemp)
		conn.Do("save") //保存
	}

	//读取当日房间状态
	data, _ := redis.Bytes(conn.Do("hget", mes.UserId, mes.Day))
	var totleMesRes message.TotalMes
	json.Unmarshal(data, &totleMesRes)

	//将读取到的当日房间状态序列化
	dayMesr, _ := json.Marshal(totleMesRes.Rooms)
	var dayMesRes message.Message
	dayMesRes.Type = message.DayRoomMesRes
	dayMesRes.Data = string(dayMesr)
	//var  dayMes = make([]message.RoomMes, roomnum)
	//json.Unmarshal(data,&dayMes)
	data, _ = json.Marshal(dayMesRes)
	//向客户端发送当日房间状态
	var temp tools.Temp
	temp.Conn = c
	err = temp.WritePkg(data)
	if err != nil {
		fmt.Println("服务器向客户端发送房间信息失败 err=", err)
		return
	}
	return
}
