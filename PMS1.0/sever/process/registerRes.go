package process

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"net"
)

func RegisterRes(data string, conn net.Conn) (err error) {
	//从链接池中获取数据库的链接
	rconn := tools.GetRedisPool()
	defer rconn.Close()
	var roomMes message.RegisterMes
	err = json.Unmarshal([]byte(data), &roomMes)
	if err != nil {
		fmt.Println("json.Unmarshal roomMes fail err=", err)
		return
	}
	_, err = rconn.Do("hmset", roomMes.UserId, "pwd", roomMes.UserPwd, "roomnum", roomMes.RoomNum)
	rconn.Do("save")
	var registerResMes message.RegistrResMes

	//向数据库写入房间信息
	if err != nil {
		registerResMes.Code = 500
		registerResMes.Error = "服务器向数据库写入房间信息失败"
		fmt.Println("服务器向数据库写入房间信息失败 err=", err)
		return
	}
	registerResMes.Code = 200

	//向数据库发送完数据后，将信息反馈给客户端
	datares, err := json.Marshal(registerResMes)
	var mes message.Message
	mes.Type = message.RegisterMesType
	mes.Data = string(datares)
	datares, err = json.Marshal(mes)
	var tf tools.Temp
	tf.Conn = conn
	err = tf.WritePkg(datares)
	if err != nil {
		fmt.Println("registerMes write fial err=", err)
		return
	}
	return
}
