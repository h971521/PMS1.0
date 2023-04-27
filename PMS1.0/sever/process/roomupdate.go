package process

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
)

// 向数据库中添加订单信息
func AddOrderRes(mesint string, c net.Conn) {
	//mesint为message.RoomUptate类型
	//获取订单信息
	var roomUpdate message.RoomUpdate
	err := json.Unmarshal([]byte(mesint), &roomUpdate)
	if err != nil {
		fmt.Println("json.Unmarshal roomupdate fail err=", err)
		return
	}
	//从链接池中获取数据库链接
	conn := tools.GetRedisPool()
	defer conn.Close()
	//将订单信息更新到数据库中
	data, err := redis.Bytes(conn.Do("hget", roomUpdate.UserId, roomUpdate.Day))
	if err != nil {
		fmt.Println("redis.Bytes(conn.Do(\"hget\", roomUpdate.UserId, roomUpdate.Day)) fail err=", err)
		return
	}
	var roomsmes message.TotalMes
	err = json.Unmarshal(data, &roomsmes)
	if err != nil {
		fmt.Println("json.Unmarshal(data, &roomsmes) fail err=", err)
		return
	}
	//更新客人信息
	roomsmes.Rooms[roomUpdate.Roomnum].CustomersName = roomUpdate.CustomersName
	roomsmes.Rooms[roomUpdate.Roomnum].CustomersTel = roomUpdate.CustomersTel
	roomsmes.Rooms[roomUpdate.Roomnum].Price = roomUpdate.Price
	roomsmes.Rooms[roomUpdate.Roomnum].CheckInTime = roomUpdate.CheckInTime
	roomsmes.TotalPrice += roomUpdate.Price //今日总计
	//当添加的订单信息不为空时，房间数量减 1
	if roomUpdate.CustomersName != "" || roomUpdate.CustomersTel != 0 || roomUpdate.Price != 0 {
		roomsmes.RoomBalance--
	}
	data, err = json.Marshal(roomsmes)
	if err != nil {
		fmt.Println("向数据库发送添加订单信息失败，err=", err)
		return
	}
	_, err = conn.Do("hset", roomUpdate.UserId, roomUpdate.Day, data)
	if err != nil {
		fmt.Println("向数据库发送添加订单信息失败，err=", err)
		return
	}

	conn.Do("save")
	var addResMes message.Message
	addResMes.Type = message.AddOrderResType
	addResMes.Data = string(data)
	data, _ = json.Marshal(addResMes)
	var tf tools.Temp
	tf.Conn = c
	tf.WritePkg(data) //向客户端返回房态信息
	return
}

// 删除订单
func DelOrderRes(mes string, c net.Conn) {
	//从链接池中获取数据库链接
	conn := tools.GetRedisPool()
	defer conn.Close()
	var delmes message.RoomUpdate
	json.Unmarshal([]byte(mes), &delmes)
	//将数据库中读取指定日期的房间状态
	data, err := redis.Bytes(conn.Do("hget", delmes.UserId, delmes.Day))
	if err != nil {
		fmt.Println("redis.Bytes(conn.Do(\"hget\", roomUpdate.UserId, roomUpdate.Day)) fail err=", err)
		return
	}
	var roomsmes message.TotalMes
	err = json.Unmarshal(data, &roomsmes)
	if err != nil {
		fmt.Println("json.Unmarshal(data, &roomsmes) fail err=", err)
		return
	}
	//更新客人信息
	roomsmes.Rooms[delmes.Roomnum].CustomersName = ""
	roomsmes.Rooms[delmes.Roomnum].CustomersTel = 0
	roomsmes.Rooms[delmes.Roomnum].CheckInTime = ""
	roomsmes.TotalPrice -= roomsmes.Rooms[delmes.Roomnum].Price //今日总计
	roomsmes.Rooms[delmes.Roomnum].Price = 0
	//剩余房量+1
	roomsmes.RoomBalance++
	data, err = json.Marshal(roomsmes)
	if err != nil {
		fmt.Println("向数据库发送添加订单信息失败，err=", err)
		return
	}
	_, err = conn.Do("hset", delmes.UserId, delmes.Day, data)
	if err != nil {
		fmt.Println("向数据库发送添加订单信息失败，err=", err)
		return
	}
	conn.Do("save")
	var DelResMes message.Message
	DelResMes.Type = message.DelOrdersResType
	if err == nil {
		DelResMes.Data = "Success"
	} else {
		DelResMes.Data = "Fail"
	}
	data, _ = json.Marshal(DelResMes)
	var tf tools.Temp
	tf.Conn = c
	tf.WritePkg(data) //向客户端返回房态信息
	return
}

// 查看指定日期指定房间的订单信息
func ReadOrderRes(mes string, c net.Conn) {
	//从链接池中获取数据库链接
	conn := tools.GetRedisPool()
	defer conn.Close()
	var readmes message.RoomUpdate
	json.Unmarshal([]byte(mes), &readmes)
	//将数据库中读取指定日期的房间状态
	data, err := redis.Bytes(conn.Do("hget", readmes.UserId, readmes.Day))
	if err != nil {
		fmt.Println("redis.Bytes(conn.Do(\"hget\", roomUpdate.UserId, roomUpdate.Day)) fail err=", err)
		return
	}
	var totalMes message.TotalMes
	json.Unmarshal(data, &totalMes)
	var ordermes message.RoomMes
	ordermes = totalMes.Rooms[readmes.Roomnum]
	fmt.Printf("%+v", ordermes)
	datares, _ := json.Marshal(ordermes)
	var readres message.Message
	readres.Type = message.ReadOrderResType
	readres.Data = string(datares)
	data, _ = json.Marshal(readres)
	var tf tools.Temp
	tf.Conn = c
	tf.WritePkg(data) //向客户端返回房态信息
	return
}

// 修改订单
func TurnOrderRes(mes string, c net.Conn) {
	var roomUpdate message.RoomUpdate
	err := json.Unmarshal([]byte(mes), &roomUpdate)
	if err != nil {
		fmt.Println("json.Unmarshal roomupdate fail err=", err)
		return
	}
	//从链接池中获取数据库链接
	conn := tools.GetRedisPool()
	defer conn.Close()
	//将订单信息更新到数据库中
	data, err := redis.Bytes(conn.Do("hget", roomUpdate.UserId, roomUpdate.Day))
	if err != nil {
		fmt.Println("redis.Bytes(conn.Do(\"hget\", roomUpdate.UserId, roomUpdate.Day)) fail err=", err)
		return
	}
	var roomsmes message.TotalMes
	err = json.Unmarshal(data, &roomsmes)
	if err != nil {
		fmt.Println("json.Unmarshal(data, &roomsmes) fail err=", err)
		return
	}
	//更新客人信息
	if roomUpdate.CustomersName != "" { //修改名字
		roomsmes.Rooms[roomUpdate.Roomnum].CustomersName = roomUpdate.CustomersName
	}
	if roomUpdate.CustomersTel != 0 {
		roomsmes.Rooms[roomUpdate.Roomnum].CustomersTel = roomUpdate.CustomersTel
	}
	if roomUpdate.Price != 0 {
		roomsmes.TotalPrice = roomsmes.TotalPrice - roomsmes.Rooms[roomUpdate.Roomnum].Price + roomUpdate.Price
		roomsmes.Rooms[roomUpdate.Roomnum].Price = roomUpdate.Price
	}
	roomsmes.Rooms[roomUpdate.Roomnum].CheckInTime = roomUpdate.CheckInTime
	data, err = json.Marshal(roomsmes)
	if err != nil {
		fmt.Println("向数据库发送添加订单信息失败，err=", err)
		return
	}
	_, err = conn.Do("hset", roomUpdate.UserId, roomUpdate.Day, data)
	if err != nil {
		fmt.Println("向数据库发送添加订单信息失败，err=", err)
		return
	}
	//var totalPrice float32 //今日总价
	//var roomBalance int    //剩余放量
	//fmt.Println(roomsmes)
	conn.Do("save")
	var addResMes message.Message
	addResMes.Type = message.TurnOrderResType
	addResMes.Data = string(data)
	data, _ = json.Marshal(addResMes)
	var tf tools.Temp
	tf.Conn = c
	tf.WritePkg(data) //向客户端返回房态信息
	return
}
